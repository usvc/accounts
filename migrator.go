package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type Migrator struct {
}

type MigratorOptions struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

var migrator Migrator

func (*Migrator) run(opt MigratorOptions) {
	connectionString := opt.User + ":" + opt.Password + "@tcp(" + opt.Host + ":" + opt.Port + ")/" + opt.Database
	db := migrator.getConnection(connectionString)
	driver := migrator.getDriver(db)
	migratorInstance := migrator.getDatabaseInstance(driver)
	migrator.migrateToLatest(migratorInstance)
	migratorInstance.Close()
}

func (*Migrator) migrateToLatest(migratorInstance *migrate.Migrate) {
	version, dirty, err := migratorInstance.Version()
	if version == 0 && err.Error() == "no migration" {
		logger.info("no migrations applied yet")
	} else {
		logger.infof("migration version: %v (dirty: %v)", version, dirty)
	}
	if dirty == true {
		logger.warnf("removing dirty migration from version %v", version)
		if err := migratorInstance.Force(int(version)); err != nil {
			logger.error(err)
		} else if err := migratorInstance.Steps(-1); err != nil {
			logger.error(err)
		}
	}

	migrationDone := false
	var migrationError *error
	for migrationDone == false && migrationError == nil {
		if err := migratorInstance.Steps(1); err != nil {
			if err.Error() == "file does not exist" {
				logger.infof("migration is up-to-date at version: %v (dirty: %v)", version, dirty)
				migrationDone = true
			} else {
				logger.errorf("migration upward failed with error: %s", err)
				migrationError = &err
			}
		} else if version, dirty, err = migratorInstance.Version(); err != nil {
			logger.error(err)
		} else {
			logger.infof("migration version: %v (dirty: %v)", version, dirty)
		}
	}

	logger.info("migration run completed")
}

func (*Migrator) getConnection(connection string) *sql.DB {
	if databaseConnection, err := sql.Open("mysql", connection); err != nil {
		logger.errorf("error while creating database connection: %s", err)
		panic(err)
	} else {
		return databaseConnection
	}
}

func (*Migrator) getDriver(databaseConnection *sql.DB) database.Driver {
	if driver, err := mysql.WithInstance(databaseConnection, &mysql.Config{}); err != nil {
		logger.errorf("error in getting driver: %s", err)
		panic(err)
	} else {
		return driver
	}
}

func (*Migrator) getDatabaseInstance(driver database.Driver) *migrate.Migrate {
	if migratorInstance, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	); err != nil {
		logger.errorf("error while creating migrator: %s", err)
		panic(err)
	} else {
		return migratorInstance
	}
}
