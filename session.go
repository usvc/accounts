package main

import (
	"database/sql"
	"fmt"
	"strings"
)

// Session represents a session where a user is logged in
type Session struct {
	ID           int    `json:"id"`
	AccountUUID  string `json:"account_uuid"`
	IPv4         string `json:"ipv4"`
	IPv6         string `json:"ipv6"`
	Source       string `json:"source"`
	Device       string `json:"device"`
	TokenRefresh string `json:"token_refresh"`
	TokenAccess  string `json:"token_access"`
	DateExpires  string `json:"date_expires"`
	DateCreated  string `json:"date_created"`
	LastModified string `json:"last_modified"`
}

// SessionError represents a logical error instead of a system one
type SessionError struct {
	Message string
	Code    string
	Data    interface{}
}

var (
	// SessionErrorOk indicates nothing went wrong
	SessionErrorOk = "E_SESSION_OK"
	// SessionErrorCreate indicates the session creation failed
	SessionErrorCreate = "E_SESSION_CREATE"
)

// Create , well, creates a new session
func (session *Session) Create(database *sql.DB) {
	logger.Infof("[session] creating new session (AccountUUID:%v,)", session.AccountUUID)

	// validation
	// this verifies the user exists
	user := User{}
	user.GetByUUID(database, session.AccountUUID)

	// insert into the database
	session.create(database)
}

// create , well, creates the session in persistent storage
func (session *Session) create(database *sql.DB) {
	// at this point, user exists, let's go
	sqlStmt, values := session.getCreateSQL()
	logger.Infof("[session] executing sql '%s'", sqlStmt)
	if stmt, err := database.Prepare(sqlStmt); err != nil {
		panic(err)
	} else if results, err := stmt.Exec(values...); err != nil {
		panic(err)
	} else if rowsAffected, err := results.RowsAffected(); err != nil {
		panic(err)
	} else if rowsAffected != 1 {
		// this basically means we screwed up somewhere
		panic(&SessionError{
			Code:    SessionErrorCreate,
			Message: "something went wrong while creating a new session, are your parameters correct?",
			Data:    session,
		})
	}
}

// getCreateSQL returns the sql statement as a string and a slice representing values
// to be fed into an .Exec
func (session *Session) getCreateSQL() (string, []interface{}) {
	var insertionsArray []string
	var valuePlaceholdersArray []string
	var valuesArray = []interface{}{
		session.AccountUUID,
	}
	sqlStmtStub := "INSERT INTO sessions (account_id %s) VALUES ((SELECT id FROM accounts WHERE uuid = ?) %s)"
	session.insertStringIntoInsertionArray(&insertionsArray, &valuesArray, session.IPv4, "ipv4")
	session.insertStringIntoInsertionArray(&insertionsArray, &valuesArray, session.IPv6, "ipv6")
	session.insertStringIntoInsertionArray(&insertionsArray, &valuesArray, session.Source, "source")
	session.insertStringIntoInsertionArray(&insertionsArray, &valuesArray, session.Device, "device")
	session.insertStringIntoInsertionArray(&insertionsArray, &valuesArray, session.TokenRefresh, "token_refresh")
	session.insertStringIntoInsertionArray(&insertionsArray, &valuesArray, session.TokenAccess, "token_access")
	session.insertStringIntoInsertionArray(&insertionsArray, &valuesArray, session.DateExpires, "date_expires")
	for i := 1; i < len(valuesArray); i++ {
		valuePlaceholdersArray = append(valuePlaceholdersArray, "?")
	}
	insertions := strings.Join(insertionsArray, ", ")
	valuePlaceholders := strings.Join(valuePlaceholdersArray, ", ")
	if len(insertionsArray) > 0 {
		insertions = fmt.Sprintf(", %s", insertions)
		valuePlaceholders = fmt.Sprintf(", %s", valuePlaceholders)
	}
	sqlStmt := fmt.Sprintf(sqlStmtStub, insertions, valuePlaceholders)
	return sqlStmt, valuesArray
}

func (session *Session) insertStringIntoInsertionArray(
	insertionsArray *[]string,
	valuesArray *[]interface{},
	value string,
	label string,
) {
	if len(value) > 0 {
		*insertionsArray = append(*insertionsArray, label)
		*valuesArray = append(*valuesArray, value)
	}
}
