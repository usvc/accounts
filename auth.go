package main

import (
	"database/sql"
	"fmt"
)

// AuthCredentials handles authentication via username/email and password
type AuthCredentials struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	TokenRefresh string `json:"token_refresh"`
	TokenAccess  string `json:"token_access"`
}

// AutheCredentialsError for authentication credential errors
type AuthError struct {
	Code    string
	Message string
	Data    interface{}
}

func (authError *AuthError) Error() string {
	return fmt.Sprintf("%s:%s", authError.Code, authError.Message)
}

var (
	// AuthErrorMissingParams indicates a missing username/email
	AuthErrorMissingParams = "E_AUTH_CREDENTIALS_MISSING_PARAMS"
	// AuthErrorInvalidParams indicates an invalid username/email/password combination
	// - we do not reveal anything more since this is sensitive information
	AuthErrorInvalidParams = "E_AUTH_CREDENTIALS_INVALID_PARAMS"
)

// Authenticate validates that the proper parameters are in place and calls
// the database function to validate the authentication
func (authCredentials *AuthCredentials) Authenticate(database *sql.DB) {
	hasUsername := len(authCredentials.Username) > 0
	hasEmail := len(authCredentials.Email) > 0
	hasPassword := len(authCredentials.Password) > 0
	if (!hasUsername && !hasEmail) || !hasPassword {
		panic(&AuthError{
			Code:    AuthErrorMissingParams,
			Message: "either 'username' or 'email', and 'password', should be specified",
		})
	}
	authCredentials.authenticate(database)
}

func (authCredentials *AuthCredentials) authenticate(database *sql.DB) {
	var sqlStmt string
	var accountIdentifier string
	sqlStmtStub := `
		SELECT sec.password
			FROM security sec
			INNER JOIN accounts acc
				ON sec.account_id = acc.id
			WHERE`
	if len(authCredentials.Username) > 0 {
		sqlStmt = fmt.Sprintf("%s acc.username = ?", sqlStmtStub)
		accountIdentifier = authCredentials.Username
	} else {
		sqlStmt = fmt.Sprintf("%s acc.email = ?", sqlStmtStub)
		accountIdentifier = authCredentials.Email
	}
	if stmt, err := database.Prepare(sqlStmt); err != nil {
		panic(err)
	} else if row := stmt.QueryRow(accountIdentifier); err != nil {
		panic(err)
	} else {
		var passwordHash sql.NullString
		if err = row.Scan(&passwordHash); err != nil {
			if err == sql.ErrNoRows {
				panic(&AuthError{
					Code:    AuthErrorInvalidParams,
					Message: "the email/username/password combination does not exist",
				})
			}
		}
		if err := utils.VerifyPasswordHash(authCredentials.Password, passwordHash.String); err != nil {
			panic(&AuthError{
				Code:    AuthErrorInvalidParams,
				Message: "the email/username/password combination does not exist",
			})
		}
	}
}
