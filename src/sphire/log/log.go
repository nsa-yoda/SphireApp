package log

import (
	"os"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

var lpx = log.New()

func Init(formatter string, environment string, log_file string) {
	switch formatter {
	case "text":
		lpx.Formatter = new(log.TextFormatter)
	case "json":
		lpx.Formatter = new(log.JSONFormatter)
	}

	// Output to stderr/file instead of stdout
	if log_file != "" {
		fmt.Println("file")
		f_ptr, _ := os.OpenFile(log_file, os.O_RDWR|os.O_APPEND, os.ModePerm)
		lpx.Out = f_ptr
		//defer f_ptr.Close()
	} else {
		fmt.Println("Stderr")
		log.SetOutput(os.Stderr)
	}

	switch environment {
	case "DEV":
		lpx.Level = log.DebugLevel
	case "STG":
		lpx.Level = log.InfoLevel
	case "PRD":
		lpx.Level = log.WarnLevel
	}
}

// Log logs a message to the defined logger
// var fields map[string]interface{} = make(map[string]interface{})
func Log(fields map[string]interface{}, message string, level string) bool {
	switch level {
	case "debug":
		lpx.WithFields(fields).Debug(message)
		return true
	case "info":
		lpx.WithFields(fields).Info(message)
		return true
	case "warn":
		lpx.WithFields(fields).Warn(message)
		return true
	case "error":
		lpx.WithFields(fields).Error(message)
		return true
	case "fatal":
		lpx.WithFields(fields).Fatal(message)
		return true
	case "panic":
		lpx.WithFields(fields).Panic(message)
		return true
	}

	return false
}