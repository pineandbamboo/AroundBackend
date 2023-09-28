package service

import (
	"around/backend"
	"around/constants"
	"around/model"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func AddUser(user *model.User) (bool, error) {
	// check whether user existed
	query := elastic.NewTermQuery("username", user.Username)
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
	if err != nil {
		return false, nil
	}

	// save User to ES
	err = backend.ESBackend.SaveToES(user, constants.USER_INDEX, user.Username)
	if err != nil {
		return false, err
	}
	fmt.Printf("User is added: %s\n", user.Username)
	return true, nil
}

// 1. search username -> password, compare password
// 2. search username + password, hit -> success verified

func CheckUser(username, password string) (bool, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("username", username))
	query.Must(elastic.NewTermQuery("password", password))
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
	if err != nil {
		return false, err
	}
	// if searchResult.Totalhits() > 0 {
	// 	return true, nil
	// }
	
	var utype model.User
	for _, item := range searchResult.Each(reflect.TypeOf(utype)) {
		u := item.(model.User)
		if u.Password == password {
			fmt.Printf("Login as %s\n", username)
			return true, nil
		}
	}
	return false, nil
}