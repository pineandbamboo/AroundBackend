package service

import (
	"around/backend"
	"around/constants"
	"around/model"
	"fmt"

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
	
}