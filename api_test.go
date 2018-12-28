package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIError(t *testing.T) {
	apiError := APIError{
		Code:    "E_CODE_TEST",
		Message: "testing",
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected an error but none was thrown")
		} else {
			if r.(*APIError).Error() != "[api] E_CODE_TEST:testing" {
				t.Errorf("error formatting is funky")
			}
		}
	}()
	panic(&apiError)
}

func TestAPIHandlerAPIError(t *testing.T) {
	apiError := &APIError{
		Code:    "_api_error",
		Message: "_api_message",
		Data:    "_api_data",
	}
	apiHandler := APIHandler(func(w http.ResponseWriter, r *http.Request) {
		panic(apiError)
	})
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://somewhere.com", nil)
	defer func() {
		results := responseWriter.Result()
		if results.StatusCode != 400 {
			t.Errorf("expected status code 400 but got %v", results.StatusCode)
		}
		body, _ := ioutil.ReadAll(results.Body)
		var apiResponse APIResponse
		err := json.Unmarshal(body, &apiResponse)
		if err != nil {
			panic(err)
		}
		if apiResponse.Code != "_api_error" {
			t.Errorf("code property expected '%v' but got '%v'", apiError.Code, apiResponse.Code)
		}
		if apiResponse.Message != "_api_message" {
			t.Errorf("message property expected '%v' but got '%v'", apiError.Message, apiResponse.Message)
		}
		if apiResponse.Data != "_api_data" {
			t.Errorf("data property expected '%v' but got '%v'", apiError.Data, apiResponse.Data)
		}
	}()
	apiHandler.ServeHTTP(responseWriter, request)
}

func TestAPIHandlerModelError(t *testing.T) {
	modelError := &ModelError{
		Code:    "_model_error",
		Message: "_model_message",
		Data:    "_model_data",
	}
	apiHandler := APIHandler(func(w http.ResponseWriter, r *http.Request) {
		panic(modelError)
	})
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://somewhere.com", nil)
	defer func() {
		results := responseWriter.Result()
		if results.StatusCode != 400 {
			t.Errorf("expected status code 400 but got %v", results.StatusCode)
		}
		body, _ := ioutil.ReadAll(results.Body)
		var apiResponse APIResponse
		err := json.Unmarshal(body, &apiResponse)
		if err != nil {
			panic(err)
		}
		if apiResponse.Code != "_model_error" {
			t.Errorf("code property expected '%v' but got '%v'", modelError.Code, apiResponse.Code)
		}
		if apiResponse.Message != "_model_message" {
			t.Errorf("message property expected '%v' but got '%v'", modelError.Message, apiResponse.Message)
		}
		if apiResponse.Data != "_model_data" {
			t.Errorf("data property expected '%v' but got '%v'", modelError.Data, apiResponse.Data)
		}
	}()
	apiHandler.ServeHTTP(responseWriter, request)
}

func TestAPIHandlerGenericError(t *testing.T) {
	errorMessage := "no error type"
	apiHandler := APIHandler(func(w http.ResponseWriter, r *http.Request) {
		panic(errorMessage)
	})
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://somewhere.com", nil)
	defer func() {
		results := responseWriter.Result()
		if results.StatusCode != 500 {
			t.Errorf("expected status code 500 but got %v", results.StatusCode)
		}
		body, _ := ioutil.ReadAll(results.Body)
		var apiResponse APIResponse
		err := json.Unmarshal(body, &apiResponse)
		if err != nil {
			panic(err)
		}
		if apiResponse.Code != APIErrorGeneric {
			t.Errorf("code property expected '%v' but got '%v'", APIErrorGeneric, apiResponse.Code)
		}
		if len(apiResponse.Message) > 0 {
			t.Errorf("message property expected '' but got '%v'", apiResponse.Message)
		}
		if apiResponse.Data != errorMessage {
			t.Errorf("data property expected '%v' but got '%v'", errorMessage, apiResponse.Data)
		}
	}()
	apiHandler.ServeHTTP(responseWriter, request)
}

func TestAPIResponseCustomHTTPCode(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	response := APIResponse{
		Code:    "_code",
		Message: "_message",
		Data:    "_data",
		status:  400,
	}
	response.send(responseWriter)
	result := responseWriter.Result()

	statusCode := result.StatusCode
	if statusCode != 400 {
		t.Errorf("http status code should not be modified")
	}

	contentType := result.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("content-type should be enforced to application/json, got '%s'", contentType)
	}
}

func TestAPIResponseCustomContentType(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	response := APIResponse{
		Code:        "_code",
		Message:     "_message",
		Data:        "_data",
		contentType: "text/plain",
	}
	response.send(responseWriter)
	result := responseWriter.Result()

	contentType := result.Header.Get("Content-Type")
	if contentType != "text/plain" {
		t.Errorf("content-type should be set to text/plain, got '%s'", contentType)
	}
}

func TestAPIResponseDataIntegrity(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	response := APIResponse{
		Code:    "_code",
		Message: "_message",
		Data:    "_data",
	}
	response.send(responseWriter)
	result := responseWriter.Result()
	statusCode := result.StatusCode
	if statusCode != 200 {
		t.Errorf("http status code was not 200")
	}

	contentType := result.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("response content-type is not application/json")
	}

	body, _ := ioutil.ReadAll(result.Body)
	var bodyInterface map[string]interface{}
	json.Unmarshal(body, &bodyInterface)
	if bodyInterface["timestamp"] == nil {
		t.Errorf("timestamp was expected but not found")
	}
	if bodyInterface["error_code"] != "_code" {
		t.Errorf("error_code was improperly passed down")
	}
	if bodyInterface["message"] != "_message" {
		t.Errorf("message was improperly passed down")
	}
	if bodyInterface["data"] != "_data" {
		t.Errorf("data was improperly passed down")
	}
}
