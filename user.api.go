package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var UserErrorOk = "E_USER_OK"
var UserErrorInvalidParameters = "E_INVALID_PARAMS"
var UserUrlStub = "/user"
var UserExtUrlStub = "/user/"

type UserAPI struct {
	router *mux.Router
}

func (userApi *UserAPI) handle(router *http.ServeMux) {
	userApi.router = mux.NewRouter()
	UserAPIGetUserByUuid(userApi.router)
	UserAPICreatetUser(userApi.router)
	router.Handle(UserUrlStub, userApi.router)
	router.Handle(UserExtUrlStub, userApi.router)
}

var userApi = UserAPI{}

func UserAPIGetUserByUuid(router *mux.Router) {
	router.HandleFunc(
		UserUrlStub+"/{uuid}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			data := user.GetByUuid(vars["uuid"])
			response := Response{
				UserErrorOk,
				"ok",
				data,
			}
			response.send(w)
		},
	).Methods("GET")
}

func UserAPICreatetUser(router *mux.Router) {
	router.HandleFunc(
		UserUrlStub,
		func(w http.ResponseWriter, r *http.Request) {
			var newUser UserNew
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &newUser)

			logger.info(newUser)
			data := user.Create(newUser)
			logger.info(data)

			response := Response{
				UserErrorOk,
				"ok",
				data,
			}
			response.send(w)
		},
	).Methods("POST")
}
