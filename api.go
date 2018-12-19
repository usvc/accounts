package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	// APIErrorGeneric is for errors where we have no idea what's going on
	APIErrorGeneric = "E_API_GENERIC"
)

// APIHandler is the wrapper around all API calls so that we can return
// a consistent schema for responses
type APIHandler func(http.ResponseWriter, *http.Request)

// ServeHTTP allows us to interface with the http.Handle
func (apiHandler APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("[api] %v", r)
			var response APIResponse
			switch t := r.(type) {
			case *SessionError:
				w.WriteHeader(400)
				response = APIResponse{
					Code:    t.Code,
					Message: t.Message,
					Data:    t.Data,
				}
			case *SecurityError:
				w.WriteHeader(400)
				response = APIResponse{
					Code:    t.Code,
					Message: t.Message,
					Data:    t.Data,
				}
			case *UserAPIError:
				w.WriteHeader(400)
				response = APIResponse{
					Code:    t.Code,
					Message: t.Message,
					Data:    t.Data,
				}
			case *UserError:
				w.WriteHeader(400)
				response = APIResponse{
					Code:    t.Code,
					Message: t.Message,
					Data:    t.Data,
				}
			default:
				w.WriteHeader(500)
				response = APIResponse{
					Code:    APIErrorGeneric,
					Message: "",
					Data:    r,
				}
			}
			response.send(w)
		}
	}()
	apiHandler(w, r)
}

// APIResponse is the schema we use for returning data to the consumer
type APIResponse struct {
	Code      string      `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}

func (response *APIResponse) send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response.Timestamp = time.Now().Format(time.RFC1123)
	body, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v", string(body))
}
