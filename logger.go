package gologger

import (
	"log"
	"os"
	"strings"
)

type LogLevelType = int

const (
	DEBUG LogLevelType = iota
	INFO
	ERROR
)

var LogLevel LogLevelType = INFO

func Init(envVarName string) {
	logLevel := os.Getenv(envVarName)
	newLogLevel := ERROR
	valid := true
	switch logLevel {
	case "DEBUG":
		newLogLevel = DEBUG
	case "INFO":
		newLogLevel = INFO
	case "ERROR":
		newLogLevel = ERROR
	default:
		valid = false
	}
	if valid {
		Info("LogLevel: %s", logLevel)
	} else {
		Info("Loglevel '%s' invalid. Falling back to Error", logLevel)
	}
	LogLevel = newLogLevel
}

func appendNewLine(msg string) string {
	if !strings.HasSuffix(msg, "\n") {
		return msg + "\n"
	}
	return msg
}

func Debug(msg string, args ...any) {
	if LogLevel == DEBUG {
		msg = appendNewLine(msg)
		log.Printf("[DEBUG] "+msg, args...)
	}
}

func Info(msg string, args ...any) {
	if LogLevel != ERROR {
		msg = appendNewLine(msg)
		log.Printf("[INFO] "+msg, args...)
	}
}

func Error(msg string, args ...any) {
	msg = appendNewLine(msg)
	log.Printf("[Error] "+msg, args...)
}

func Fatal(msg string, args ...any) {
	msg = appendNewLine(msg)
	log.Fatalf("[Fatal] "+msg, args...)
}
