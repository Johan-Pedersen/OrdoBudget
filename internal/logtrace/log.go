package logging

import (
	"log"
	"os"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

// Initialize the loggers during package initialization
func init() {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an info message
func Info(msg string) {
	infoLog.Println(msg)
}

// Error logs an error message
func Error(msg string) {
	errorLog.Println(msg)
}

