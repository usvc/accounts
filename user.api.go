package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	// UserAPIErrorOK indicates its all okay
	UserAPIErrorOK = "E_USER_API_OK"
	// UserAPIErrorDeleteOK indicates deletion is okay
	UserAPIErrorDeleteOK = "E_USER_API_DELETE_OK"
	// UserAPIErrorCreateOk indicates user creation is okay
	UserAPIErrorCreateOk = "E_USER_API_CREATE_OK"
	// UserAPIErrorCreateGeneric represents a generic user creation error
	UserAPIErrorCreateGeneric = "E_USER_API_CREATE_GENERIC"
	// UserAPIErrorGetOK indicates getting a single user is okay
	UserAPIErrorGetOK = "E_USER_API_GET_OK"
	// UserAPIErrorQueryOk indicates user querying is okay
	UserAPIErrorQueryOk = "E_USER_API_QUERY_OK"
	// UserAPIErrorQueryInvalidParameters indicates user querying is okay
	UserAPIErrorQueryInvalidParameters = "E_USER_API_QUERY_INVALID_PARAMETERS"
	// UserAPIErrorUpateOk indicates user updating is okay
	UserAPIErrorUpateOk = "E_USER_API_UPDATE_OK"
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
	userApi.handleQueryUsers(userApi.router)
	userApi.handleCreateUser(userApi.router)
	userApi.handleDeleteUserByUUID(userApi.router)
	userApi.handleUpdateUser(userApi.router)
	router.Handle(UserAPIUrlStub, userApi.router)
	router.Handle(UserAPIExtURLStub, userApi.router)
}

func (userApi *UserAPI) handleUpdateUser(router *mux.Router) {
	router.Handle(
		UserAPIExtURLStub+"{uuid}",
		APIHandler(func(w http.ResponseWriter, r *http.Request) {
			params := mux.Vars(r)
			var userData User
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &userData)
			userData.UUID = params["uuid"]
			user.UpdateByUUID(db.Get(), &userData)
			response := APIResponse{
				Code:    UserAPIErrorUpateOk,
				Message: "",
				Data:    userData,
			}
			response.send(w)
		}),
	).Methods("PATCH")
}

func (userApi *UserAPI) handleQueryUsers(router *mux.Router) {
	router.Handle(
		UserAPIUrlStub,
		APIHandler(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			startIndex := 0
			dataLimit := 10
			if len(query["start_at"]) > 0 {
				_startIndex, err := strconv.Atoi(query["start_at"][0])
				if err != nil {
					panic(&APIError{
						Code:    UserAPIErrorQueryInvalidParameters,
						Message: err.Error(),
						Data:    query["start_at"][0],
					})
				} else if _startIndex > 0 {
					startIndex = _startIndex
				}
			}
			if len(query["limit"]) > 0 {
				_dataLimit, err := strconv.Atoi(query["limit"][0])
				if err != nil {
					panic(&APIError{
						Code:    UserAPIErrorQueryInvalidParameters,
						Message: err.Error(),
						Data:    query["limit"][0],
					})
				} else if _dataLimit > 0 {
					dataLimit = _dataLimit
				}
			}
			data := user.Query(db.Get(), uint(startIndex), uint(dataLimit))
			response := APIResponse{
				Code:    UserAPIErrorQueryOk,
				Message: "ok",
				Data:    data,
			}
			response.send(w)
		}),
	).Methods("GET")
}

func (userApi *UserAPI) handleDeleteUserByUUID(router *mux.Router) {
	router.Handle(
		UserAPIUrlStub+"/{uuid}",
		APIHandler(func(w http.ResponseWriter, r *http.Request) {
			params := mux.Vars(r)
			user.DeleteByUUID(db.Get(), params["uuid"])
			response := APIResponse{
				Code:    UserAPIErrorDeleteOK,
				Message: "ok",
				Data:    map[string]interface{}{"uuid": params["uuid"]},
			}
			response.send(w)
		}),
	).Methods("DELETE")
}

func (userApi *UserAPI) handleGetUserByUUID(router *mux.Router) {
	router.Handle(
		UserAPIUrlStub+"/{uuid}",
		APIHandler(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			data := user.GetByUUID(db.Get(), vars["uuid"])
			response := APIResponse{
				Code:    UserAPIErrorGetOK,
				Message: "ok",
				Data:    data,
			}
			response.send(w)
		}),
	).Methods("GET")
}

func (userApi *UserAPI) handleCreateUser(router *mux.Router) {
	router.Handle(
		UserAPIUrlStub,
		APIHandler(func(w http.ResponseWriter, r *http.Request) {
			var newUser UserNew
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &newUser)
			data := user.Create(db.Get(), newUser)

			response := APIResponse{
				Code:    UserAPIErrorCreateOk,
				Message: "ok",
				Data:    data,
			}
			response.send(w)
		}),
	).Methods("POST")
}
