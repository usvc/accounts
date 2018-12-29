package main

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
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

func TestAuthenticateUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Errorf("unexpected error while creating database stub connection: '%s'", err)
	}
	password := "password"
	hashedPassword, err := utils.CreatePasswordHash(password)
	if err != nil {
		t.Errorf("unexpected error while creating a mock password hash: '%s'", err)
	}
	mock.ExpectPrepare("^SELECT sec.password FROM security sec .+ WHERE acc.username = ?")
	mock.
		ExpectQuery("^SELECT sec.password FROM security sec .+ WHERE acc.username = ?").
		WillReturnRows(sqlmock.
			NewRows([]string{"password"}).
			AddRow(hashedPassword),
		)
	mock.ExpectCommit()
	auth := AuthCredentials{
		Username: "username",
		Password: password,
	}
	auth.Authenticate(db)
}

func TestAuthenticateEmail(t *testing.T) {
	password := "password"
	expectedQueryRegex := "^SELECT sec.password FROM security sec .+ WHERE acc.email = ?"
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Errorf("unexpected error while creating database stub connection: '%s'", err)
	}
	hashedPassword, err := utils.CreatePasswordHash(password)
	if err != nil {
		t.Errorf("unexpected error while creating a mock password hash: '%s'", err)
	}
	mock.ExpectPrepare(expectedQueryRegex)
	mock.
		ExpectQuery(expectedQueryRegex).
		WillReturnRows(sqlmock.
			NewRows([]string{"password"}).
			AddRow(hashedPassword),
		)
	mock.ExpectCommit()
	auth := AuthCredentials{
		Email:    "email@dmain.com",
		Password: password,
	}
	auth.Authenticate(db)
}
