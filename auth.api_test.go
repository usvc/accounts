package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type TestSuiteAuthAPI struct {
	suite.Suite
	handler  *http.ServeMux
	server   *httptest.Server
	instance AuthAPI
	db       *sql.DB
	dbMock   sqlmock.Sqlmock
}

func TestRunTestSuiteAuthAPI(t *testing.T) {
	suite.Run(t, new(TestSuiteAuthAPI))
}

func (suite *TestSuiteAuthAPI) SetupTest() {
	var err error
	var mock sqlmock.Sqlmock
	suite.db, mock, err = sqlmock.New()
	if err != nil {
		suite.T().Errorf("u")
	}
	suite.dbMock = mock
	suite.handler = http.NewServeMux()
	suite.instance = AuthAPI{}
	suite.instance.Handle(suite.handler, suite.db)
	suite.server = httptest.NewServer(suite.handler)
}

func (suite *TestSuiteAuthAPI) TeardownTest() {
	suite.db.Close()
	suite.server.Close()
}

func (suite *TestSuiteAuthAPI) TestEnsureCompatibility() {
	defer func() {
		assert.Nil(suite.T(), recover())
	}()
}

func (suite *TestSuiteAuthAPI) TestLoginWithUsernameCredentials() {
	var apiResponse APIResponse
	testRequestBody := fmt.Sprintf(`{"username":"test","password":"%s"}`, authTestPassword)
	testRequestURL := strings.Trim(suite.server.URL, "/") + AuthAPIExtURLStub + "credentials"
	expectedQueryRegex := "^SELECT sec.password FROM security sec .+ WHERE acc.username = ?"
	hashedPassword, err := utils.CreatePasswordHash(authTestPassword)
	assert.Nil(suite.T(), err)

	// setup mocks & expectations
	suite.dbMock.ExpectPrepare(expectedQueryRegex)
	suite.dbMock.ExpectQuery(expectedQueryRegex).
		WillReturnRows(sqlmock.
			NewRows([]string{"password"}).
			AddRow(hashedPassword),
		)
	suite.dbMock.ExpectCommit()

	// issue test request
	response, err := suite.server.Client().
		Post(testRequestURL,
			"application/json",
			bytes.NewBuffer([]byte(testRequestBody)),
		)
	assert.Nil(suite.T(), err)
	responseBody, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), json.Unmarshal(responseBody, &apiResponse))
	assert.Equal(suite.T(), apiResponse.Code, "E_AUTH_API_OK")
}

func (suite *TestSuiteAuthAPI) TestLoginWithEmailCredentials() {
	var apiResponse APIResponse
	testRequestBody := fmt.Sprintf(`{"email":"email@domain.com","password":"%s"}`, authTestPassword)
	testRequestURL := strings.Trim(suite.server.URL, "/") + AuthAPIExtURLStub + "credentials"
	expectedQueryRegex := "^SELECT sec.password FROM security sec .+ WHERE acc.email = ?"
	hashedPassword, err := utils.CreatePasswordHash(authTestPassword)
	assert.Nil(suite.T(), err)

	// setup mocks & expectations
	suite.dbMock.ExpectPrepare(expectedQueryRegex)
	suite.dbMock.ExpectQuery(expectedQueryRegex).
		WillReturnRows(sqlmock.
			NewRows([]string{"password"}).
			AddRow(hashedPassword),
		)
	suite.dbMock.ExpectCommit()

	// issue test request
	response, err := suite.server.Client().
		Post(testRequestURL,
			"application/json",
			bytes.NewBuffer([]byte(testRequestBody)),
		)
	assert.Nil(suite.T(), err)
	responseBody, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), json.Unmarshal(responseBody, &apiResponse))
	assert.Equal(suite.T(), apiResponse.Code, "E_AUTH_API_OK")
}
