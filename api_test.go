package main

import (
	"encoding/json"
	"io/ioutil"
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
