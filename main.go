package main

import (
	"fmt"
	"strings"
	"time"
)

const applicationStartImage = `
                   _                        _       
 _ _ ___ _ _ ___  / |__ ___ ___ ___ _ _ ___| |_ ___ 
| | |_ -| | |  _|/ / .'|  _|  _| . | | |   |  _|_ -|
|___|___|\_/|___|_/|__,|___|___|___|___|_|_|_| |___|
`

func main() {
	defer func() {
		if r := recover(); r != nil {
			logger.withStackf("%s", r)
		}
	}()
	fmt.Println(applicationStartImage)
	logger.init(config.Environment)
	config.init()
	logger.infof("usvc/accounts started in %s environment at %s", strings.ToUpper(config.Environment), time.Now().Format(time.RFC1123Z))
	if config.IsMigration {
		logger.info("performing migration...")
		migrator.run(&migratorOptions{
			Host:     config.DatabaseHost,
			Port:     config.DatabasePort,
			Database: config.DatabaseDB,
			User:     config.DatabaseUser,
			Password: config.DatabasePassword,
		})
	} else {
		logger.info("starting service...")
		server.init(&serverOptions{
			Interface: config.Interface,
			Port:      config.Port,
		})
	}
}
