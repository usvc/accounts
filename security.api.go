package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// SecurityAPI module for interfacing with the Security module
type SecurityAPI struct {
	router *mux.Router
	model  *Security
}

var (
	// SecurityAPIErrorPasswordChangeOk indicates password change went well
	SecurityAPIErrorPasswordChangeOk = "E_SECURITY_PASSWORD_CHANGE_OK"
	// SecurityAPIUrlStub for '/security' endpoints
	SecurityAPIUrlStub = "/security"
	// SecurityAPIExtURLStub for '/security/*' endpoints
	SecurityAPIExtURLStub = "/security/"
)

// Handle adds support for a :router to serve paths at `/security/*`
func (securityApi *SecurityAPI) Handle(router *http.ServeMux) {
	securityApi.model = &Security{}
	securityApi.router = mux.NewRouter()
	securityApi.updatePassword(securityApi.router)
	router.Handle(SecurityAPIUrlStub, securityApi.router)
	router.Handle(SecurityAPIExtURLStub, securityApi.router)
}

func (securityApi *SecurityAPI) updatePassword(router *mux.Router) {
	router.Handle(
		SecurityAPIExtURLStub+"{account_uuid}",
		APIHandler(func(w http.ResponseWriter, r *http.Request) {
			params := mux.Vars(r)
			accountUUID := params["account_uuid"]
			var passwordData Security
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &passwordData)
			securityApi.model.UpdatePasswordByUUID(db.Get(), passwordData.Password, accountUUID)
			response := APIResponse{
				Code:    SecurityAPIErrorPasswordChangeOk,
				Message: "ok",
				Data:    nil,
			}
			response.send(w)
		}),
	).Methods("PATCH")
}
