package main

import (
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
)

// Logger represents a dependency inversion for logging
type Logger struct {
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
)

// Init initialises the logger
func (logger *Logger) Init(opts *LoggerOptions) {
	logger.options = opts
	logrus.SetReportCaller(opts.EnableSourceMap)
	switch strings.ToLower(opts.Level) {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	}
	switch strings.ToLower(opts.Format) {
	case LoggerFormatJSON:
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: opts.EnablePrettyPrint,
		})
	case LoggerFormatText:
		fallthrough
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:               opts.EnablePrettyPrint,
			EnvironmentOverrideColors: true, // for DWDs (developers with disabilities)
			QuoteEmptyFields:          true,
		})
	}
	logrus.Infof("[logger] level: '%s', format: '%s', pretty print: '%v', sourcemap: '%v'", opts.Level, opts.Format, opts.EnablePrettyPrint, opts.EnableSourceMap)
}

// Debug level of logging
func (*Logger) Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Debugf logs at the Debug level with parameters
func (*Logger) Debugf(message string, args ...interface{}) {
	logrus.Debugf(message, args...)
}

// Trace level of logging
func (*Logger) Trace(args ...interface{}) {
	logrus.Trace(args...)
}

// Tracef logs at the Trace level with parameters
func (*Logger) Tracef(message string, args ...interface{}) {
	logrus.Tracef(message, args...)
}

// Info level of logging
func (*Logger) Info(args ...interface{}) {
	logrus.Info(args...)
}

// Infof logs at the Info level with parameters
func (*Logger) Infof(message string, args ...interface{}) {
	logrus.Infof(message, args...)
}

// Warn level of logging
func (*Logger) Warn(args ...interface{}) {
	logrus.Warn(args...)
}

// Warnf logs at the Warn level with parameters
func (*Logger) Warnf(message string, args ...interface{}) {
	logrus.Warnf(message, args...)
}

// Error level of logging
func (*Logger) Error(args ...interface{}) {
	logrus.Error(args...)
}

// Errorf logs at the Error level with parameters
func (*Logger) Errorf(message string, args ...interface{}) {
	logrus.Errorf(message, args...)
}

// Panic level of logging
func (*Logger) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

// Panicf logs at the Panic level with parameters
func (*Logger) Panicf(message string, args ...interface{}) {
	logrus.Panicf(message, args...)
}

// Fatal level of logging
func (*Logger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// Fatalf logs at the fatal level with parameters
func (*Logger) Fatalf(message string, args ...interface{}) {
	logrus.Fatalf(message, args...)
}

// WithStack logs at the error level
func (*Logger) WithStack(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"stack": strings.Split(string(debug.Stack()), "\n"),
	}).Error(args...)
}

// WithStackf logs at the error level with parameters
func (*Logger) WithStackf(message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"stack": strings.Split(string(debug.Stack()), "\n"),
	}).Errorf(message, args...)
}

// initialise global logger for ease of us
var logger = Logger{}
