package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// User is used for returning user data
type User struct {
	Uuid     string `json:"uuid"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// UserNew is used for incoming users to be created
type UserNew struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
}

// UserError represents a logical error instead of a system one
type UserError struct {
	Message string
	Code    string
	Data    interface{}
}

// Error implementation
func (userError *UserError) Error() string {
	return fmt.Sprintf("%v:%v", userError.Code, userError.Message)
}

var UserErrorCreateDuplicateEntry = "E_USER_CREATE_DUPLICATE"
var UserErrorCreateGeneric = "E_USER_CREATE_GENERIC"
var UserErrorCreateMissingParameters = "E_USER_CREATE_MISSING_PARAMS"
var UserErrorCreateInvalidEmail = "E_USER_CREATE_INVALID_EMAIL"
var UserErrorCreateInvalidPassword = "E_USER_CREATE_INVALID_PASSWORD"

var userStatementsPrepared = false

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
		panic(&UserError{
			Code:    UserErrorCreateMissingParameters,
			Message: "missing 'email' parameter",
		})
	} else if len(newUser.Password) == 0 {
		panic(&UserError{
			Code:    UserErrorCreateMissingParameters,
			Message: "missing 'password' parameter",
		})
	}

	// validate parameters
	if err := utils.ValidateEmail(newUser.Email); err != nil {
		panic(&UserError{
			Code:    err.Error(),
			Message: "",
			Data:    map[string]interface{}{"email": newUser.Email},
		})
	} else if err := utils.ValidatePassword(newUser.Password); err != nil {
		panic(&UserError{
			Code:    err.Error(),
			Message: "",
			Data:    map[string]interface{}{},
		})
	}

	// prepare what to put into the database
	if passwordHash, err := utils.CreatePasswordHash(newUser.Password); err != nil {
		panic(err)
	} else {
		newUser.PasswordHash = passwordHash
	}

	// put it into the database
	userID := user.insertUser(newUser.Email)
	user.insertSecurity(userID, newUser.PasswordHash)
	userRow := user.getUserByID(userID)

	logger.infof("[user] created user '%s'", userRow.Uuid)

	return userRow
}

func (*User) getUserByID(accountId int64) User {
	stmt, err := db.Get().Prepare("SELECT uuid, email, username FROM accounts WHERE id = ?")
	if err != nil {
		panic(err)
	}
	row := stmt.QueryRow(accountId)
	if err != nil {
		panic(err)
	}
	var uuid sql.NullString
	var email sql.NullString
	var username sql.NullString
	err = row.Scan(&uuid, &email, &username)
	if err != nil {
		panic(err)
	}
	return User{
		Uuid:     uuid.String,
		Email:    email.String,
		Username: username.String,
	}
}

func (*User) insertUser(email string) int64 {
	logger.info("[user] adding account data...")
	stmt, err := db.Get().Prepare("INSERT INTO accounts (email) VALUES (?)")
	if err != nil {
		panic(err)
	}
	output, err := stmt.Exec(email)
	if err != nil {
		logger.errorf("[user] %v", err)
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			panic(&UserError{
				Code:    UserErrorCreateDuplicateEntry,
				Message: "the user already exists",
				Data:    err,
			})
		default:
			panic(UserErrorCreateGeneric)
		}
	}
	lastInsertID, err := output.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertID
}

func (*User) insertSecurity(accountId int64, passwordHash string) {
	logger.info("[user] adding account security...")
	stmt, err := db.Get().Prepare("INSERT INTO security (account_id, password) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(accountId, passwordHash)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if rowsAffected != 1 {
		panic("[user] expected 1 row to be affected but none were")
	}
}
