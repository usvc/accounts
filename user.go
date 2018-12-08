package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uuid     string `json:"uuid"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserNew struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
}

var UserErrorCreateOk = "E_USER_CREATE_OK"
var UserErrorCreateMissingParameters = "E_USER_CREATE_MISSING_PARAMS"
var UserErrorCreateInvalidEmail = "E_USER_CREATE_INVALID_EMAIL"
var UserErrorCreateInvalidPassword = "E_USER_CREATE_INVALID_PASSWORD"

var user = User{}

func (user *User) GetByUuid(uuid string) User {
	logger.info("getUser()")
	return User{
		Uuid:     uuid,
		Email:    "todo@todo.com",
		Username: "username",
	}
}

func (user *User) Create(newUser UserNew) User {
	// check for missing parameters
	if len(newUser.Email) == 0 {
		panic(UserErrorCreateMissingParameters)
	} else if len(newUser.Password) == 0 {
		panic(UserErrorCreateMissingParameters)
	}

	// validate parameters
	if err := utils.ValidateEmail(newUser.Email); err != nil {
		panic(UserErrorCreateInvalidEmail)
	} else if err := utils.ValidatePassword(newUser.Password); err != nil {
		panic(UserErrorCreateInvalidPassword)
	}

	// prepare what to put into the database
	if passwordHash, err := utils.CreatePasswordHash(newUser.Password); err != nil {
		panic(err)
	} else {
		newUser.PasswordHash = passwordHash
	}

	// put it into the database
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		"user",
		"password",
		"database",
		"3306",
		"database",
	))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		logger.info("it works!")
	}

	stmtCreateAccount, err := db.Prepare("INSERT INTO accounts (email) VALUES (?)")
	if err != nil {
		panic(err)
	}
	result, err := stmtCreateAccount.Exec(newUser.Email)
	if err != nil {
		panic(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	stmtCreateSecurity, err := db.Prepare("INSERT INTO security (account_id, password) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	result, err = stmtCreateSecurity.Exec(lastInsertID, newUser.PasswordHash)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if rowsAffected != 1 {
		panic("expected 1 row to be affected but none were")
	}
	stmtGetUser, err := db.Prepare("SELECT uuid FROM accounts WHERE id = ?")
	if err != nil {
		panic(err)
	}
	row := stmtGetUser.QueryRow(lastInsertID)
	if err != nil {
		panic(err)
	}
	var uuid string
	err = row.Scan(&uuid)
	if err != nil {
		panic(err)
	}
	logger.info(uuid)

	return User{
		Uuid:  uuid,
		Email: newUser.Email,
	}
}
