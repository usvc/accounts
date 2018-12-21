package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// AuthAPIErrorOK indicates all is fine
	AuthAPIErrorOK = "E_AUTH_API_OK"
	// AuthAPIURLStub provisions for /auth
	AuthAPIURLStub = "/auth"
	// AuthAPIExtURLStub provisions for /auth/*
	AuthAPIExtURLStub = "/auth/"
)

// AuthAPI handles the /auth and /auth/* endpoints
type AuthAPI struct {
	router *mux.Router
}

// Handle is the handler for the AuthAPI module
func (authAPI *AuthAPI) Handle(router *http.ServeMux) {
	authAPI.router = mux.NewRouter()
	authAPI.handleCredentialsLogin()
	router.Handle(AuthAPIURLStub, authAPI.router)
	router.Handle(AuthAPIExtURLStub, authAPI.router)
}

func (authAPI *AuthAPI) handleCredentialsLogin() {
	authAPI.router.Handle(
		AuthAPIExtURLStub+"credentials",
		APIHandler(func(w http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			var authCredentials AuthCredentials
			json.Unmarshal(body, &authCredentials)
			authCredentials.Authenticate(db.Get())
			response := APIResponse{
				Code:    AuthAPIErrorOK,
				Message: "ok",
			}
			response.send(w)
		}),
	).Methods("POST")
}
