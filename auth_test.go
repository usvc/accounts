package main

import (
	"testing"
)

func TestAuthenticateValidationNoEmailAndUsername(t *testing.T) {
	auth := AuthCredentials{
		Password: "password",
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected a panic but none happened")
		}
	}()
	auth.Authenticate(nil)
}

func TestAuthenticateValidationOnlyUsername(t *testing.T) {
	auth := AuthCredentials{
		Username: "username",
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected a panic but none happened")
		}
	}()
	auth.Authenticate(nil)
}

func TestAuthenticateValidationOnlyEmail(t *testing.T) {
	auth := AuthCredentials{
		Email: "email@domain.com",
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected a panic but none happened")
		}
	}()
	auth.Authenticate(nil)
}
