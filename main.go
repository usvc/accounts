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
			logger.WithStackf("%s", r)
		}
	}()
	config := Configuration{}
	config.Init()
	fmt.Println(applicationStartImage)
	logger.Init(&LoggerOptions{
		EnableSourceMap:   config.LogSourceMap,
		EnablePrettyPrint: config.LogPrettyPrint,
		Format:            config.LogFormat,
		Level:             config.LogLevel,
	})
	logger.Info("[main] initialising database connection...")
	db.Init(&DatabaseConnectionOptions{
		Host:                      config.DatabaseHost,
		Port:                      config.DatabasePort,
		Database:                  config.DatabaseDB,
		User:                      config.DatabaseUser,
		Password:                  config.DatabasePassword,
		ConnectionRetryAttempts:   3,
		ConnectionRetryIntervalMs: 5000,
	})
	logger.Infof("[main] usvc/accounts started in %s environment at %s", strings.ToUpper(config.Environment), time.Now().Format(time.RFC1123Z))
	if config.IsMigration {
		logger.Info("[main] performing migration...")
		migrator := Migrator{}
		migrator.run(&MigratorConnectionOptions{
			Host:                      config.DatabaseHost,
			Port:                      config.DatabasePort,
			Database:                  config.DatabaseDB,
			User:                      config.DatabaseUser,
			Password:                  config.DatabasePassword,
			ConnectionRetryAttempts:   3,
			ConnectionRetryIntervalMs: 5000,
		})
	} else {
		logger.Info("[main] starting server...")
		server := Server{}
		server.init(&ServerOptions{
			Interface: config.Interface,
			Port:      config.Port,
		})
	}
}
