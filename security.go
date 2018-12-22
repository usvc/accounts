package main

import (
	"database/sql"
)

// Security module for handling security related account data
type Security struct {
	AccountUUID    string `json:"account_uuid"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashed_password"`
}

// UpdatePasswordByUUID sets the password of the user identified by account UUID :accountUUID
// to the :password.
func (security *Security) UpdatePasswordByUUID(database *sql.DB) {
	if err := utils.ValidatePassword(security.Password); err != nil {
		panic(&ModelError{
			Code:    err.(*ValidationError).Code,
			Message: err.(*ValidationError).Message,
			Data:    map[string]interface{}{}, // reveal nothing, it's the password (:
		})
	}
	hashedPassword, err := utils.CreatePasswordHash(security.Password)
	if err != nil {
		panic(err)
	}
	security.HashedPassword = hashedPassword
	security.updatePasswordByUUID(database)
}

func (security *Security) updatePasswordByUUID(database *sql.DB) {
	sqlStmt := "UPDATE security AS s INNER JOIN accounts AS a ON s.account_id = a.id SET s.password = ? WHERE a.uuid = ?"
	logger.Infof("[security] executing sql '%s'", sqlStmt)
	stmt, err := database.Prepare(sqlStmt)
	if err != nil {
		panic(err)
	}
	results, err := stmt.Exec(security.HashedPassword, security.AccountUUID)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		panic(err)
	} else if rowsAffected == 0 {
		user := &User{}
		user.GetByUUID(database, security.AccountUUID)
	}
}
