package main

import (
	"flag"
	"github.com/spf13/viper"
)

type Configuration struct {
	Environment string
	Interface string
	Port string
	IsMigration bool
	DatabaseHost string
	DatabasePort string
	DatabaseDB string
	DatabaseUser string
	DatabasePassword string
}

var config = Configuration{}

// Init initialises the configuration module
func (config *Configuration) Init() {
	config.configureEnvironment()
	config.configureEntrypoint()
}

// configureEnvironment consumes environment variables into config variables
func (*Configuration) configureEnvironment() {
	viper.SetDefault("environment", "development")
	viper.SetDefault("interface", "0.0.0.0")
	viper.SetDefault("port", "3000")
	viper.SetDefault("db_host", "database")
	viper.SetDefault("db_port", "3306")
	viper.SetDefault("db_database", "database")
	viper.SetDefault("db_user", "user")
	viper.SetDefault("db_password", "password")
	viper.AutomaticEnv()
	config = Configuration{
		Environment: viper.GetString("environment"),
		Interface: viper.GetString("interface"),
		Port: viper.GetString("port"),
		DatabaseHost: viper.GetString("db_host"),
		DatabasePort: viper.GetString("db_port"),
		DatabaseDB: viper.GetString("db_database"),
		DatabaseUser: viper.GetString("db_user"),
		DatabasePassword: viper.GetString("db_password"),
	}
}

// configureEntrypoint consumes flags and sets the run mode flags
func (*Configuration) configureEntrypoint() {
	isMigration := flag.Bool("migrate", false, "indicates whether application should attempt migration")
	flag.Parse()
	config.IsMigration = *isMigration
}
