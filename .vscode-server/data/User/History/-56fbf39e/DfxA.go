package handler

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options {
        ValidationKeyGetter : func(token *jwt.Token) (interface{}, error) {
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