package main

import "fmt"

// ModelError represents an error originating from the
// model layer
type ModelError struct {
	Code    string
	Message string
	Data    interface{}
}

func (modelError *ModelError) Error() string {
	return fmt.Sprintf("[user.api] %v:%v", modelError.Code, modelError.Message)
}
