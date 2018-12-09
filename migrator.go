package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type Migrator struct {
	db       *sql.DB
	driver   database.Driver
	instance *migrate.Migrate
	options  *MigratorConnectionOptions
}

type MigratorConnectionOptions struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type migratorConstSettings struct {
	maxRetries      uint
	retryIntervalMs uint
}

var migratorSettings = migratorConstSettings{
	maxRetries:      3,
	retryIntervalMs: 5000,
}

var migrator = Migrator{}

func (migrator *Migrator) run(opts *MigratorConnectionOptions) {
	migrator.options = opts
	connectionString := migrator.getConnectionString()
	migrator.db = migrator.getConnection(connectionString)
	migrator.driver = migrator.getDriver()
	migrator.instance = migrator.getDatabaseInstance()
	migrator.migrateToLatest()
	migrator.instance.Close()

	logger.info("[migrator] migration run completed")
}

func (migrator *Migrator) getConnectionString() string {
	logger.infof("[migrator] connecting to database '%v' at <%v:%v> with user '%v'",
		migrator.options.Database,
		migrator.options.Host,
		migrator.options.Port,
		migrator.options.User,
	)
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		migrator.options.User,
		migrator.options.Password,
		migrator.options.Host,
		migrator.options.Port,
		migrator.options.Database,
	)
}

func (migrator *Migrator) migrateToLatest() {
	version, dirty, err := migrator.instance.Version()
	if version == 0 && err.Error() == "no migration" {
		logger.info("[migrator] no migrations applied yet")
	} else {
		logger.infof("[migrator] migration version: %v (dirty: %v)", version, dirty)
	}
	if dirty == true {
		migrator.rollbackDirtyMigration(version)
	}

	migrationDone := false
	for migrationDone == false {
		migrationDone = migrator.migrateUpwards()
	}
}

func (migrator *Migrator) rollbackDirtyMigration(version uint) {
	logger.warnf("[migrator] removing dirty migration from version %v", version)
	if err := migrator.instance.Force(int(version)); err != nil {
		logger.error(err)
	} else if err := migrator.instance.Steps(-1); err != nil {
		logger.error(err)
		panic(err)
	}
}

func (migrator *Migrator) migrateUpwards() bool {
	version, dirty, err := migrator.instance.Version()
	if err != nil {
		logger.error(err)
	}
	if err := migrator.instance.Steps(1); err != nil {
		if err.Error() == "file does not exist" {
			logger.infof("[migrator] migration is up-to-date at version: %v (dirty: %v)", version, dirty)
			return true
		} else {
			logger.errorf("[migrator] migration upward failed with error: %s", err)
			panic(err)
		}
	} else if version, dirty, err = migrator.instance.Version(); err != nil {
		logger.error(err)
	} else {
		logger.infof("[migrator] migration version now at %v (dirty: %v)", version, dirty)
	}
	return false
}

func (*Migrator) getConnection(connection string) *sql.DB {
	if databaseConnection, err := sql.Open("mysql", connection); err != nil {
		logger.errorf("[migrator] error while creating database connection: %s", err)
		panic(err)
	} else {
		return databaseConnection
	}
}

func (migrator *Migrator) getDriver() database.Driver {
	var currentTry uint
	var driver database.Driver
	var err error
	for currentTry = 0; currentTry < migratorSettings.maxRetries; currentTry++ {
		if driver, err = mysql.WithInstance(migrator.db, &mysql.Config{}); err != nil {
			logger.errorf("[migrator] failed to get driver (current try: %v/%v), error: %s", currentTry, migratorSettings.maxRetries, err)
			time.Sleep(time.Duration(migratorSettings.retryIntervalMs) * time.Millisecond)
		} else {
			return driver
		}
	}
	logger.errorf("[migrator] error in getting driver: %s", err)
	panic(err)
}

func (migrator *Migrator) getDatabaseInstance() *migrate.Migrate {
	if instance, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		migrator.driver,
	); err != nil {
		logger.errorf("[migrator] error while creating migrator: %s", err)
		panic(err)
	} else {
		return instance
	}
}
