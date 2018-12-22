package main

import (
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
