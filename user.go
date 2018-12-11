package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// User is used for returning user data
type User struct {
	Uuid         string `json:"uuid"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	DateCreated  string `json:"date_created"`
	LastModified string `json:"last_modified"`
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

// Error implementation for UserError
func (userError *UserError) Error() string {
	return fmt.Sprintf("%v:%v", userError.Code, userError.Message)
}

var (
	// UserErrorNotFound for errors where user is not found
	UserErrorNotFound = "E_USER_NOT_FOUND"
	// UserErrorCreateGeneric for generic user creation errors
	UserErrorCreateGeneric = "E_USER_CREATE_GENERIC"
	// UserErrorCreateDuplicateEntry for errors where a duplicate user is found
	UserErrorCreateDuplicateEntry = "E_USER_CREATE_DUPLICATE"
	// UserErrorCreateMissingParameters for missing parameters
	UserErrorCreateMissingParameters = "E_USER_CREATE_MISSING_PARAMS"
	// UserErrorCreateInvalidEmail for invalid emails
	UserErrorCreateInvalidEmail = "E_USER_CREATE_INVALID_EMAIL"
	// UserErrorCreateInvalidPassword for invalid passwords
	UserErrorCreateInvalidPassword = "E_USER_CREATE_INVALID_PASSWORD"
)

var userStatementsPrepared = false

var user = User{}

/*
-------------------------------------------------------------------------------
USER CREATION
#create #registration #new #account
-------------------------------------------------------------------------------
*/

func (user *User) Create(database *sql.DB, newUser UserNew) *User {
	logger.Infof("[user] creating user with email '%s'", newUser.Email)

	// validate parameters
	if len(newUser.Email) == 0 {
		panic(&UserError{
			Code:    UserErrorCreateMissingParameters,
			Message: "missing 'email' parameter",
		})
	} else if err := utils.ValidateEmail(newUser.Email); err != nil {
		panic(&UserError{
			Code:    err.(*ValidationError).Code,
			Message: err.(*ValidationError).Message,
			Data:    map[string]interface{}{"email": newUser.Email},
		})
	} else if len(newUser.Password) == 0 {
		panic(&UserError{
			Code:    UserErrorCreateMissingParameters,
			Message: "missing 'password' parameter",
		})
	} else if err := utils.ValidatePassword(newUser.Password); err != nil {
		panic(&UserError{
			Code:    err.(*ValidationError).Code,
			Message: err.(*ValidationError).Message,
			Data:    map[string]interface{}{}, // reveal nothing, it's the password (:
		})
	}

	userRow := user.create(database, &newUser)

	logger.Infof("[user] created user with email '%s' - uuid is '%s'", newUser.Email, userRow.Uuid)

	return userRow
}

// create executes the database operations for Create()
func (*User) create(database *sql.DB, newUser *UserNew) *User {
	if passwordHash, err := utils.CreatePasswordHash(newUser.Password); err != nil {
		panic(err)
	} else {
		newUser.PasswordHash = passwordHash
	}

	userID := user.createAccount(database, newUser.Email)
	user.createSecurity(database, userID, newUser.PasswordHash)
	return user.getByID(database, userID)
}

// createAccount adds a new row to the `accounts` table, returns the ID of the
// newly created user
func (*User) createAccount(database *sql.DB, email string) int64 {
	logger.Info("[user] adding account data...")
	stmt, err := database.Prepare("INSERT INTO accounts (email) VALUES (?)")
	if err != nil {
		panic(err)
	}
	output, err := stmt.Exec(email)
	if err != nil {
		logger.Errorf("[user] %v", err)
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			panic(&UserError{
				Code:    UserErrorCreateDuplicateEntry,
				Message: "the user already exists",
				Data:    map[string]interface{}{"email": email},
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

// createSecurity adds the security details of the user with ID :id
func (*User) createSecurity(database *sql.DB, id int64, passwordHash string) {
	logger.Info("[user] adding account security...")
	stmt, err := database.Prepare("INSERT INTO security (account_id, password) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(id, passwordHash)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	} else if rowsAffected != 1 {
		panic("[user] expected 1 row to be affected but none were")
	}
}

/*
-------------------------------------------------------------------------------
USER QUERY
#read #account
-------------------------------------------------------------------------------
*/

func (user *User) Query(database *sql.DB, startIndex uint, limit uint) *[]User {
	logger.Infof("[user] querying %v users starting from index %v...", limit, startIndex)
	stmt, err := database.Prepare("SELECT uuid, email, username, date_created, last_modified FROM accounts LIMIT ?,?")
	if err != nil {
		panic(err)
	}
	var users []User
	rows, err := stmt.Query(startIndex, limit)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var uuid sql.NullString
		var email sql.NullString
		var username sql.NullString
		var dateCreated sql.NullString
		var lastModified sql.NullString
		err := rows.Scan(&uuid, &email, &username, &dateCreated, &lastModified)
		if err != nil {
			panic(err)
		}
		users = append(users, User{
			Uuid:         uuid.String,
			Email:        email.String,
			Username:     username.String,
			DateCreated:  dateCreated.String,
			LastModified: lastModified.String,
		})
	}
	logger.Infof("[user] queried %v users (requested for %v)", len(users), limit)
	return &users
}

// GetByUUID retrieves a user with UUID :uuid
func (user *User) GetByUUID(database *sql.DB, uuid string) *User {
	logger.Infof("[user] getting user with UUID '%v'", uuid)

	if len(uuid) == 0 {
		panic(&UserError{
			Code:    UserErrorCreateMissingParameters,
			Message: "missing 'uuid' parameter",
		})
	}

	userRow := user.getByUUID(database, uuid)

	logger.Infof("[user] retrieved user with UUID '%v'", uuid)

	return userRow
}

// getByUUID executes the database operations for GetByUUID()
func (*User) getByUUID(database *sql.DB, uuid string) *User {
	stmt, err := database.Prepare("SELECT email, username, date_created, last_modified FROM accounts WHERE uuid = ?")
	if err != nil {
		panic(err)
	}
	row := stmt.QueryRow(uuid)
	if err != nil {
		logger.Errorf("[user] %v", err)
		panic(err)
	}
	var email sql.NullString
	var username sql.NullString
	var dateCreated sql.NullString
	var lastModified sql.NullString
	err = row.Scan(&email, &username, &dateCreated, &lastModified)
	if err != nil {
		if err == sql.ErrNoRows {
			panic(&UserError{
				Code:    UserErrorNotFound,
				Message: fmt.Sprintf("the user identified by %s does not exist", uuid),
			})
		} else {
			panic(err)
		}
	}
	return &User{
		Uuid:         uuid,
		Email:        email.String,
		Username:     username.String,
		DateCreated:  dateCreated.String,
		LastModified: lastModified.String,
	}
}

// getById executes the database operations to retrieve a user identified by
// the ID :id
func (*User) getByID(database *sql.DB, id int64) *User {
	stmt, err := database.Prepare("SELECT uuid, email, username, date_created, last_modified FROM accounts WHERE id = ?")
	if err != nil {
		panic(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		panic(err)
	}
	var uuid sql.NullString
	var email sql.NullString
	var username sql.NullString
	var dateCreated sql.NullString
	var lastModified sql.NullString
	err = row.Scan(&uuid, &email, &username, &dateCreated, &lastModified)
	if err != nil {
		panic(err)
	}
	return &User{
		Uuid:         uuid.String,
		Email:        email.String,
		Username:     username.String,
		DateCreated:  dateCreated.String,
		LastModified: lastModified.String,
	}
}

/*
-------------------------------------------------------------------------------
USER REMOVAL
#delete #remove #account #user
-------------------------------------------------------------------------------
*/

// DeleteByUUID removes the user identified by :uuid using the database
// connection :database
func (user *User) DeleteByUUID(database *sql.DB, uuid string) {
	logger.Infof("[user] removing user with UUID '%v'", uuid)

	if len(uuid) == 0 {
		panic(&UserError{
			Code:    UserErrorCreateMissingParameters,
			Message: "missing 'uuid' parameter",
		})
	}

	user.deleteByUUID(database, uuid)

	logger.Infof("[user] removed user with UUID '%v'", uuid)
}

// deleteByUUID defines the database operations for removing a user identified
// by :uuid
func (*User) deleteByUUID(database *sql.DB, uuid string) {
	stmt, err := database.Prepare("DELETE FROM accounts WHERE uuid = ?")
	if err != nil {
		panic(err)
	}
	exec, err := stmt.Exec(uuid)
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		panic(err)
	} else if rowsAffected == 0 {
		panic(&UserError{
			Code:    UserErrorNotFound,
			Message: fmt.Sprintf("user with uuid '%s' could not be found", uuid),
		})
	}
}
