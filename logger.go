package main

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// LoggerConfig provides a way to configure the logger
type LoggerConfig struct {
	options *LoggerOptions
}

// LoggerOptions provides a configuration class for the Logger
type LoggerOptions struct {
	Level             string
	Format            string
	EnablePrettyPrint bool
	EnableSourceMap   bool
}

var (
	// LoggerFormatJSON defines the JSON type of logging
	LoggerFormatJSON = "json"
	// LoggerFormatText defines the Text type of logging
	LoggerFormatText = "text"
	// global logger
	logger = logrus.New()
	// global logger configurer
	loggerConfig = LoggerConfig{}
)

// Init initialises the logger
func (*LoggerConfig) Init(opts *LoggerOptions) {
	logger.SetReportCaller(opts.EnableSourceMap)
	switch strings.ToLower(opts.Level) {
	case "trace":
		logger.Level = logrus.TraceLevel
	case "debug":
		logger.Level = logrus.DebugLevel
	case "info":
		logger.Level = logrus.InfoLevel
	case "warn":
		logger.Level = logrus.WarnLevel
	case "error":
		logger.Level = logrus.ErrorLevel
	case "fatal":
		logger.Level = logrus.FatalLevel
	case "panic":
		logger.Level = logrus.PanicLevel
	}
	switch strings.ToLower(opts.Format) {
	case LoggerFormatJSON:
		logger.Formatter = &logrus.JSONFormatter{
			PrettyPrint: opts.EnablePrettyPrint,
		}
	case LoggerFormatText:
		fallthrough
	default:
		logger.Formatter = &logrus.TextFormatter{
			ForceColors:               opts.EnablePrettyPrint,
			EnvironmentOverrideColors: true, // for DWDs (developers with disabilities)
			QuoteEmptyFields:          true,
			FullTimestamp:             true,
			TimestampFormat:           "Jan 2/3:4 PM MST", // refer to https://golang.org/src/time/format.go
		}
	}
	logger.Infof("[logger] configured with {level: '%s', format: '%s', pretty print: '%v', sourcemap: '%v'}", opts.Level, opts.Format, opts.EnablePrettyPrint, opts.EnableSourceMap)
}
