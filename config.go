package main

import (
	"flag"
	"github.com/spf13/viper"
)

// Configuration defines the possible environment variables we will be consuming
type Configuration struct {
	Environment string
	Interface string
	Port string
	IsMigration bool
	LogFormat string
	LogLevel string
	LogPrettyPrint bool
	LogSourceMap bool
	DatabaseHost string
	DatabasePort string
	DatabaseDB string
	DatabaseUser string
	DatabasePassword string
}

// Init initialises the configuration module
func (config *Configuration) Init() {
	config.configureEnvironment()
	config.configureEntrypoint()
}

// configureEnvironment consumes environment variables into config variables
func (config *Configuration) configureEnvironment() {
	viper.SetDefault("environment", "development")
	viper.SetDefault("interface", "0.0.0.0")
	viper.SetDefault("port", "3000")
	viper.SetDefault("log_format", "text")
	viper.SetDefault("log_level", "trace")
	viper.SetDefault("log_source_map", true)
	viper.SetDefault("log_pretty_print", true)
	viper.SetDefault("db_host", "database")
	viper.SetDefault("db_port", "3306")
	viper.SetDefault("db_database", "database")
	viper.SetDefault("db_user", "user")
	viper.SetDefault("db_password", "password")
	viper.AutomaticEnv()
	config.Environment = viper.GetString("environment")
	config.Interface = viper.GetString("interface")
	config.Port = viper.GetString("port")
	config.LogFormat = viper.GetString("log_format")
	config.LogLevel = viper.GetString("log_level")
	config.LogSourceMap = viper.GetBool("log_source_map")
	config.LogPrettyPrint = viper.GetBool("log_pretty_print")
	config.DatabaseHost = viper.GetString("db_host")
	config.DatabasePort = viper.GetString("db_port")
	config.DatabaseDB = viper.GetString("db_database")
	config.DatabaseUser = viper.GetString("db_user")
	config.DatabasePassword = viper.GetString("db_password")
}

// configureEntrypoint consumes flags and sets the run mode flags
func (config *Configuration) configureEntrypoint() {
	isMigration := flag.Bool("migrate", false, "indicates whether application should attempt migration")
	flag.Parse()
	config.IsMigration = *isMigration
}
