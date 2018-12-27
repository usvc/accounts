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
	// APIErrorTodo is for to-be implemented endpoints
	APIErrorTodo = "E_API_TODO"
)

// APIError is a wrapper error type for errors originating from
// the API layer
type APIError struct {
	Code    string
	Message string
	Data    interface{}
}

func (apiError *APIError) Error() string {
	return fmt.Sprintf("[api] %v:%v", apiError.Code, apiError.Message)
}

// APIHandler is the wrapper around all API calls so that we can return
// a consistent schema for responses
type APIHandler func(http.ResponseWriter, *http.Request)

// ServeHTTP allows us to interface with the http.Handle
func (apiHandler APIHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	start := time.Now()
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("[api] %v", r)
			var response APIResponse
			switch t := r.(type) {
			case *ModelError:
				response = APIResponse{
					Code:    t.Code,
					Message: t.Message,
					Data:    t.Data,
					status:  400,
				}
			case *APIError:
				response = APIResponse{
					Code:    t.Code,
					Message: t.Message,
					Data:    t.Data,
					status:  400,
				}
			default:
				response = APIResponse{
					Code:    APIErrorGeneric,
					Message: "",
					Data:    r,
					status:  500,
				}
			}
			response.send(res)
		}
		logger.Info(map[string]interface{}{
			"proto":        req.Proto,
			"method":       req.Method,
			"path":         req.URL.Path,
			"hostname":     req.Host,
			"remoteAddr":   req.RemoteAddr,
			"responseTime": time.Since(start).Seconds() * 1000,
			"userAgent":    req.Header.Get("user-agent"),
		})
	}()
	apiHandler(res, req)
}

// APIResponse is the schema we use for returning data to the consumer
type APIResponse struct {
	Code        string      `json:"error_code"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Timestamp   string      `json:"timestamp"`
	contentType string
	status      int
}

func (response *APIResponse) send(w http.ResponseWriter) {
	if len(response.contentType) > 0 {
		w.Header().Set("Content-Type", response.contentType)
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
	if response.status > 0 {
		w.WriteHeader(response.status)
	} else {
		w.WriteHeader(200)
	}
	response.Timestamp = time.Now().Format(time.RFC1123)
	body, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v", string(body))
}
