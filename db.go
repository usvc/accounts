package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Database is the primary module for interacting with
// persistent data
type Database struct {
	inst    *sql.DB
	options *DatabaseConnectionOptions
}

// DatabaseConnectionOptions provides a class for configuring
// the database module
type DatabaseConnectionOptions struct {
	Host                      string
	Port                      string
	Database                  string
	User                      string
	Password                  string
	ConnectionRetryIntervalMs time.Duration
	ConnectionRetryAttempts   uint
}

var db Database

// Init initialises the database module
func (database *Database) Init(opts *DatabaseConnectionOptions) {
	database.options = opts
	database.createConnection()
	database.validateConnection()
}

// Get returns the database instance
func (database *Database) Get() *sql.DB {
	return database.inst
}

func (database *Database) createConnection() {
	logger.Infof("[db] creating connection with '%s:%s' using user '%s'...", database.options.Host, database.options.Port, database.options.User)
	instance, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		database.options.User,
		database.options.Password,
		database.options.Host,
		database.options.Port,
		database.options.Database,
	))
	if err != nil {
		panic(err)
	}
	database.inst = instance
}

func (database *Database) validateConnection() {
	err := errors.New("uninitialised")
	var connectionAttempt uint
	for connectionAttempt = 1; err != nil && connectionAttempt <= database.options.ConnectionRetryAttempts; connectionAttempt++ {
		logger.Infof("[db] pinging database (%v/%v attempts) to validate connection...", connectionAttempt, database.options.ConnectionRetryAttempts)
		err = database.inst.Ping()
		if err != nil {
			logger.Infof("[db] database ping failed, waiting %v milliseconds before trying again...", database.options.ConnectionRetryIntervalMs)
			time.Sleep(database.options.ConnectionRetryIntervalMs * time.Millisecond)
		} else {
			logger.Info("[db] ping succeeded, proceeding...")
		}
	}
	if err != nil {
		panic(err)
	}
}
