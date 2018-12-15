package main

import (
	"database/sql"
)

// Security module for handling security related account data
type Security struct {
	Password string `json:"password"`
}

// SecurityError represents a logical error instead of a system one
type SecurityError struct {
	Message string
	Code    string
	Data    interface{}
}

// UpdatePasswordByUUID sets the password of the user identified by account UUID :accountUUID
// to the :password.
func (security *Security) UpdatePasswordByUUID(database *sql.DB, password string, accountUUID string) {
	if err := utils.ValidatePassword(password); err != nil {
		panic(&SecurityError{
			Code:    err.(*ValidationError).Code,
			Message: err.(*ValidationError).Message,
			Data:    map[string]interface{}{}, // reveal nothing, it's the password (:
		})
	}
	hashedPassword, err := utils.CreatePasswordHash(password)
	if err != nil {
		panic(err)
	}
	security.updatePasswordByUUID(database, &User{}, hashedPassword, accountUUID)
}

func (security *Security) updatePasswordByUUID(database *sql.DB, user *User, hashedPassword string, accountUUID string) {
	sqlStmt := "UPDATE security AS s INNER JOIN accounts AS a ON s.account_id = a.id SET s.password = ? WHERE a.uuid = ?"
	logger.Infof("[security] executing sql '%s'", sqlStmt)
	stmt, err := database.Prepare(sqlStmt)
	if err != nil {
		panic(err)
	}
	results, err := stmt.Exec(hashedPassword, accountUUID)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		panic(err)
	} else if rowsAffected == 0 {
		user.GetByUUID(database, accountUUID)
	}
}
