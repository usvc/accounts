package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var UserApiErrorOK = "E_USER_API_OK"
var UserApiErrorDeleteOK = "E_USER_API_ERROR_DELETE_OK"
var UserApiErrorInvalidParameters = "E_USER_API_INVALID_PARAMS"
var UserApiErrorCreateOk = "E_USER_API_CREATE_OK"
var UserApiErrorCreateGeneric = "E_USER_API_CREATE_GENERIC"
var UserApiUrlStub = "/user"
var UserApiExtUrlStub = "/user/"

type UserAPI struct {
	router *mux.Router
}

func (userApi *UserAPI) handle(router *http.ServeMux) {
	userApi.router = mux.NewRouter()
	UserAPIGetUserByUuid(userApi.router)
	UserAPICreateUser(userApi.router)
	UserAPIDeleteUserByUuid(userApi.router)
	router.Handle(UserApiUrlStub, userApi.router)
	router.Handle(UserApiExtUrlStub, userApi.router)
}

var userApi = UserAPI{}

func UserAPIDeleteUserByUuid(router *mux.Router) {
	router.HandleFunc(
		UserApiUrlStub+"/{uuid}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			user.DeleteByUuid(vars["uuid"])
			response := Response{
				UserApiErrorDeleteOK,
				"ok",
				map[string]interface{}{"uuid": vars["uuid"]},
			}
			response.send(w)
		},
	).Methods("DELETE")
}

func UserAPIGetUserByUuid(router *mux.Router) {
	router.HandleFunc(
		UserApiUrlStub+"/{uuid}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			data := user.GetByUuid(vars["uuid"])
			response := Response{
				UserApiErrorOK,
				"ok",
				data,
			}
			response.send(w)
		},
	).Methods("GET")
}

func UserAPICreateUser(router *mux.Router) {
	router.HandleFunc(
		UserApiUrlStub,
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					var response Response
					switch t := r.(type) {
					case *UserError:
						w.WriteHeader(400)
						response = Response{
							t.Code,
							t.Message,
							t.Data,
						}
					default:
						w.WriteHeader(500)
						response = Response{
							UserApiErrorCreateGeneric,
							"",
							r,
						}
					}
					response.send(w)
				}
			}()
			var newUser UserNew
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &newUser)
			data := user.Create(newUser)

			response := Response{
				UserApiErrorCreateOk,
				"ok",
				data,
			}
			response.send(w)
		},
	).Methods("POST")
}
