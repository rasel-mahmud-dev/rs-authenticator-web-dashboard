package utils

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct{}

const (
	INFO  = "INFO"
	DEBUG = "DEBUG"
	WARN  = "WARN"
	ERROR = "ERROR"
)

var (
	once   sync.Once
	logger *Logger
)

var LoggerInstance = logger

// NewLogger returns a singleton logger instance
func NewLogger() *Logger {
	once.Do(func() {
		logger = &Logger{}
	})
	return logger
}

func init() {
	NewLogger()
}

// log is a generic function to log messages based on the level
func (l *Logger) log(level string, msg string, args ...interface{}) {
	logMessage := fmt.Sprintf(msg, args...)
	switch level {
	case INFO:
		log.Printf("[INFO] %s", logMessage)
	case DEBUG:
		log.Printf("[DEBUG] %s", logMessage)
	case WARN:
		log.Printf("[WARN] %s", logMessage)
	case ERROR:
		log.Printf("[ERROR] %s", logMessage)
	default:
		log.Printf("[INFO] %s", logMessage)
	}
}

// Info logs an informational message
func (l *Logger) Info(msg string, args ...interface{}) {
	l.log(INFO, msg, args...)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log(DEBUG, msg, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log(WARN, msg, args...)
}

// Error logs an error message
func (l *Logger) Error(msg string, args ...interface{}) {
	l.log(ERROR, msg, args...)
}

// PrintToFile logs a message to a file (optional)
func (l *Logger) PrintToFile(msg string) {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Writing log message to file
	log.SetOutput(logFile)
	log.Println(msg)
	log.SetOutput(os.Stdout) // Reset output to stdout
}
