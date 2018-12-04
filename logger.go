package main

import (
	"runtime/debug"
	"strings"

	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

type Logger struct{}

// Init initialises the logger
func (*Logger) init(environment string) {
	switch strings.ToLower(environment) {
	case "production":
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "development":
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&joonix.FluentdFormatter{})
	}
}

// Debug logs a 'debug' level message
func (*Logger) debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Debugf logs a 'debug' level message with parameters
func (*Logger) debugf(message string, args ...interface{}) {
	logrus.Debugf(message, args...)
}

func (*Logger) trace(args ...interface{}) {
	logrus.Trace(args...)
}

func (*Logger) tracef(message string, args ...interface{}) {
	logrus.Tracef(message, args...)
}

func (*Logger) info(args ...interface{}) {
	logrus.Info(args...)
}

func (*Logger) infof(message string, args ...interface{}) {
	logrus.Infof(message, args...)
}

func (*Logger) warn(args ...interface{}) {
	logrus.Warn(args...)
}

func (*Logger) warnf(message string, args ...interface{}) {
	logrus.Warnf(message, args...)
}

func (*Logger) error(args ...interface{}) {
	logrus.Error(args...)
}

func (*Logger) errorf(message string, args ...interface{}) {
	logrus.Errorf(message, args...)
}

func (*Logger) panic(args ...interface{}) {
	logrus.Panic(args...)
}

func (*Logger) panicf(message string, args ...interface{}) {
	logrus.Panicf(message, args...)
}

func (*Logger) fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (*Logger) fatalf(message string, args ...interface{}) {
	logrus.Fatalf(message, args...)
}

func (*Logger) withStack(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"stack": strings.Split(string(debug.Stack()), "\n"),
	}).Error(args...)
}

func (*Logger) withStackf(message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"stack": strings.Split(string(debug.Stack()), "\n"),
	}).Errorf(message, args...)
}

var logger = Logger{}
