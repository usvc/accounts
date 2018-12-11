package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// UserAPIErrorOK indicates its all okay
	UserAPIErrorOK = "E_USER_API_OK"
	// UserAPIErrorDeleteOK indicates deletion is okay
	UserAPIErrorDeleteOK = "E_USER_API_ERROR_DELETE_OK"
	// UserAPIErrorCreateOk indicates user creation is okay
	UserAPIErrorCreateOk = "E_USER_API_CREATE_OK"
	// UserAPIErrorCreateGeneric represents a generic user creation error
	UserAPIErrorCreateGeneric = "E_USER_API_CREATE_GENERIC"
	// UserAPIUrlStub is the base stub
	UserAPIUrlStub = "/user"
	// UserAPIExtURLStub is the extended stub
	UserAPIExtURLStub = "/user/"
)

// UserAPI is the controller layer
type UserAPI struct {
	router *mux.Router
}

// Handle takes in a router and provisions it with the user API
func (userApi *UserAPI) Handle(router *http.ServeMux) {
	userApi.router = mux.NewRouter()
	userApi.handleGetUserByUUID(userApi.router)
	userApi.handleCreateUser(userApi.router)
	userApi.handleDeleteUserByUUID(userApi.router)
	router.Handle(UserAPIUrlStub, userApi.router)
	router.Handle(UserAPIExtURLStub, userApi.router)
}

var userApi = UserAPI{}

func (userApi *UserAPI) handleDeleteUserByUUID(router *mux.Router) {
	router.HandleFunc(
		UserAPIUrlStub+"/{uuid}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			user.DeleteByUUID(db.Get(), vars["uuid"])
			response := Response{
				UserAPIErrorDeleteOK,
				"ok",
				map[string]interface{}{"uuid": vars["uuid"]},
			}
			response.send(w)
		},
	).Methods("DELETE")
}

func (userApi *UserAPI) handleGetUserByUUID(router *mux.Router) {
	router.HandleFunc(
		UserAPIUrlStub+"/{uuid}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			data := user.GetByUUID(db.Get(), vars["uuid"])
			response := Response{
				UserAPIErrorOK,
				"ok",
				data,
			}
			response.send(w)
		},
	).Methods("GET")
}

func (userApi *UserAPI) handleCreateUser(router *mux.Router) {
	router.HandleFunc(
		UserAPIUrlStub,
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
							UserAPIErrorCreateGeneric,
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
			data := user.Create(db.Get(), newUser)

			response := Response{
				UserAPIErrorCreateOk,
				"ok",
				data,
			}
			response.send(w)
		},
	).Methods("POST")
}
