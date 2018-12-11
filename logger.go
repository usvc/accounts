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
func (*Logger) Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Debugf logs a 'debug' level message with parameters
func (*Logger) Debugf(message string, args ...interface{}) {
	logrus.Debugf(message, args...)
}

func (*Logger) Trace(args ...interface{}) {
	logrus.Trace(args...)
}

func (*Logger) Tracef(message string, args ...interface{}) {
	logrus.Tracef(message, args...)
}

func (*Logger) Info(args ...interface{}) {
	logrus.Info(args...)
}

func (*Logger) Infof(message string, args ...interface{}) {
	logrus.Infof(message, args...)
}

func (*Logger) Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func (*Logger) Warnf(message string, args ...interface{}) {
	logrus.Warnf(message, args...)
}

func (*Logger) Error(args ...interface{}) {
	logrus.Error(args...)
}

func (*Logger) Errorf(message string, args ...interface{}) {
	logrus.Errorf(message, args...)
}

func (*Logger) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func (*Logger) Panicf(message string, args ...interface{}) {
	logrus.Panicf(message, args...)
}

func (*Logger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (*Logger) Fatalf(message string, args ...interface{}) {
	logrus.Fatalf(message, args...)
}

func (*Logger) WithStack(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"stack": strings.Split(string(debug.Stack()), "\n"),
	}).Error(args...)
}

func (*Logger) WithStackf(message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"stack": strings.Split(string(debug.Stack()), "\n"),
	}).Errorf(message, args...)
}

var logger = Logger{}
