package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const (
	authTestPassword = "password"
)

// TestSuiteAuth is the entry struct for testing the Auth component
type TestSuiteAuth struct {
	suite.Suite
	authCredentials AuthCredentials
}

// TestRunTestAuth runs the test suite
func TestRunTestAuth(t *testing.T) {
	suite.Run(t, new(TestSuiteAuth))
}

func (suite *TestSuiteAuth) TestNoEmailOrUsername() {
	suite.authCredentials = AuthCredentials{
		Password: authTestPassword,
	}
	defer func() {
		assert.NotNil(suite.T(), recover())
	}()
	suite.authCredentials.Authenticate(nil)
}

func (suite *TestSuiteAuth) TestOnlyUsername() {
	suite.authCredentials = AuthCredentials{
		Username: "username",
	}
	defer func() {
		assert.NotNil(suite.T(), recover())
	}()
	suite.authCredentials.Authenticate(nil)
}

func (suite *TestSuiteAuth) TestOnlyEmail() {
	suite.authCredentials = AuthCredentials{
		Email: "email@domain.com",
	}
	defer func() {
		assert.NotNil(suite.T(), recover())
	}()
	suite.authCredentials.Authenticate(nil)
}

func TestAuthenticateUsername(t *testing.T) {
	expectedQueryRegex := "^SELECT sec.password FROM security sec .+ WHERE acc.username = ?"
	hashedPassword, err := utils.CreatePasswordHash(authTestPassword)
	if err != nil {
		t.Errorf("unexpected error while creating a mock password hash: '%s'", err)
	}
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Errorf("unexpected error while creating database stub connection: '%s'", err)
	}
	mock.ExpectPrepare(expectedQueryRegex)
	mock.ExpectQuery(expectedQueryRegex).
		WillReturnRows(sqlmock.
			NewRows([]string{"password"}).
			AddRow(hashedPassword),
		)
	mock.ExpectCommit()
	auth := AuthCredentials{
		Username: "username",
		Password: authTestPassword,
	}
	auth.Authenticate(db)
}

func TestAuthenticateEmail(t *testing.T) {
	expectedQueryRegex := "^SELECT sec.password FROM security sec .+ WHERE acc.email = ?"
	hashedPassword, err := utils.CreatePasswordHash(authTestPassword)
	if err != nil {
		t.Errorf("unexpected error while creating a mock password hash: '%s'", err)
	}
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Errorf("unexpected error while creating database stub connection: '%s'", err)
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
		Email:    "email@domain.com",
		Password: authTestPassword,
	}
	auth.Authenticate(db)
}
