package middlewares

import (
	"DevBookAPI/src/authentication"
	"DevBookAPI/src/responses"
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authentication.ValidateToken(r)
		if err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}

		nextFunc(w, r)
	}
}
