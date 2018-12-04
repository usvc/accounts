package main

import (
	"flag"
	"github.com/spf13/viper"
)

type Configuration struct {
	Interface string
	Port string
	LogLevel string
	IsMigration bool
	DatabaseHost string
	DatabasePort string
	DatabaseDB string
	DatabaseUser string
	DatabasePassword string
}

var config Configuration

func (self *Configuration) init() {
	self.configureEnvironment()
	self.configureEntrypoint()
}

func (*Configuration) configureEnvironment() {
	viper.SetDefault("port", "3000")
	viper.SetDefault("log_level", "trace")
	viper.SetDefault("interface", "0.0.0.0")
	viper.SetDefault("db_host", "database")
	viper.SetDefault("db_port", "3306")
	viper.SetDefault("db_database", "database")
	viper.SetDefault("db_user", "user")
	viper.SetDefault("db_password", "password")
	viper.AutomaticEnv()
	config = Configuration{
		LogLevel: viper.GetString("log_level"),
		Port: viper.GetString("port"),
		Interface: viper.GetString("interface"),
		DatabaseHost: viper.GetString("db_host"),
		DatabasePort: viper.GetString("db_port"),
		DatabaseDB: viper.GetString("db_database"),
		DatabaseUser: viper.GetString("db_user"),
		DatabasePassword: viper.GetString("db_password"),
	}
}

func (*Configuration) configureEntrypoint() {
	isMigration := flag.Bool("migrate", false, "indicates whether application should attempt migration")
	flag.Parse()
	config.IsMigration = *isMigration
}
