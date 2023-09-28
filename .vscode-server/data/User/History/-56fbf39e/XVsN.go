package handler

import (
	"net/http"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

    jwtMiddleware := jwtMiddleware.New(jwtmiddleware.Options {
        ValidationKeyGetter : func(token *jwt.token) (interface{}, error) {
            return []byte(mySigningKey), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })


    router := mux.NewRouter()
    router.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST")
    router.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET")
    return router
}



// process
// handle
// request