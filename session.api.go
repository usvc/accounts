package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// SessionAPI is the API handler for the /session/* endpoint
type SessionAPI struct {
	router *mux.Router
}

// SessionAPIError is the standardised error for this module
type SessionAPIError struct {
	Code    string
	Message string
	Data    interface{}
}

var (
	// SessionAPIErrorOK indicates a call to /session/* is ok
	SessionAPIErrorOK = "E_SESSIONS_API_OK"
	// SessionAPIErrorCreateOK indicates call to POST /session is ok
	SessionAPIErrorCreateOK = "E_SESSIONS_API_CREATE_OK"
	// SessionAPIUrlStub represents the endpoint at /session
	SessionAPIUrlStub = "/session"
	// SessionAPIExtURLStub represents the endpoint at /session/*
	SessionAPIExtURLStub = "/session/"
)

// Error implements the error type
func (sessionAPIError *SessionAPIError) Error() string {
	return fmt.Sprintf("[sessions.api] %v:%v", sessionAPIError.Code, sessionAPIError.Message)
}

// Handle takes in a ServeMux and provisions it with the sessions API
func (sessionAPI *SessionAPI) Handle(router *http.ServeMux) {
	sessionAPI.router = mux.NewRouter()
	sessionAPI.handleCreateSessions()
	router.Handle(SessionAPIUrlStub, sessionAPI.router)
	router.Handle(SessionAPIExtURLStub, sessionAPI.router)
}

func (sessionAPI *SessionAPI) handleCreateSessions() {
	sessionAPI.router.Handle(
		SessionAPIUrlStub,
		APIHandler(func(w http.ResponseWriter, r *http.Request) {
			body, _ := ioutil.ReadAll(r.Body)
			var session Session
			json.Unmarshal(body, &session)
			session.Create(db.Get())
			response := APIResponse{
				Code:    SessionAPIErrorCreateOK,
				Message: "ok",
			}
			response.send(w)
		}),
	).Methods("POST")
}
