package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type APIHandler func(http.ResponseWriter, *http.Request)

func (this APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("[api] %v", r)
			var response APIResponse
			switch t := r.(type) {
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
					Code:    UserAPIErrorCreateGeneric,
					Message: "",
					Data:    r,
				}
			}
			response.send(w)
		}
	}()
	this(w, r)
}

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
