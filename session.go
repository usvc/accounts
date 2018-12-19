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
)

// Create
func (session *Session) Create(database *sql.DB) {
	logger.Infof("[session] creating new session (AccountUUID:%v,)", session.AccountUUID)

	// validation

	// insert into the database
	session.create(database)
}

func (session *Session) create(database *sql.DB) {
	var insertionsArray []string
	var valuePlaceholdersArray []string
	var values = []interface{}{
		session.AccountUUID,
	}
	sqlStmtStub := "INSERT INTO sessions (account_id %s) VALUES ((SELECT id FROM accounts WHERE uuid = ?) %s)"
	if len(session.IPv4) > 0 {
		insertionsArray = append(insertionsArray, "ipv4")
		values = append(values, session.IPv4)
	}
	if len(session.IPv6) > 0 {
		insertionsArray = append(insertionsArray, "ipv6")
		values = append(values, session.IPv6)
	}
	if len(session.Source) > 0 {
		insertionsArray = append(insertionsArray, "source")
		values = append(values, session.Source)
	}
	if len(session.Device) > 0 {
		insertionsArray = append(insertionsArray, "device")
		values = append(values, session.Device)
	}
	for i := 1; i < len(values); i++ {
		valuePlaceholdersArray = append(valuePlaceholdersArray, "?")
	}
	insertions := strings.Join(insertionsArray, ", ")
	valuePlaceholders := strings.Join(valuePlaceholdersArray, ", ")
	if len(insertionsArray) > 0 {
		insertions = fmt.Sprintf(", %s", insertions)
		valuePlaceholders = fmt.Sprintf(", %s", valuePlaceholders)
	}
	sqlStmt := fmt.Sprintf(sqlStmtStub, insertions, valuePlaceholders)
	logger.Infof("[session] executing sql '%s'", sqlStmt)
	stmt, err := database.Prepare(sqlStmt)
	if err != nil {
		panic(err)
	}
	logger.Info(values)
	results, err := stmt.Exec(values...)
	if err != nil {
		panic(err)
	}
	logger.Info(stmt)
	logger.Info(insertions)
	logger.Info(results)
}
