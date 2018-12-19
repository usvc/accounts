package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

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
	// SessionAPIErrorOK
	SessionAPIErrorOK = "E_SESSIONS_API_OK"
	// SessionAPIErrorNotImplemented
	SessionAPIErrorNotImplemented = "E_SESSIONS_API_TODO"
	// SessionAPIUrlStub
	SessionAPIUrlStub = "/session"
	// SessionAPIExtUrlStub
	SessionAPIExtUrlStub = "/session/"
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
	router.Handle(SessionAPIExtUrlStub, sessionAPI.router)
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
				Code:    SessionAPIErrorNotImplemented,
				Message: "",
				Data:    session,
			}
			response.send(w)
		}),
	).Methods("POST")
}
