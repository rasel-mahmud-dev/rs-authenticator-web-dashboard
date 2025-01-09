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

func NewLogger() *Logger {
	once.Do(func() {
		logger = &Logger{}
	})

	return logger
}

func init() {
	NewLogger()
}

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

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log(INFO, msg, args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log(DEBUG, msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log(WARN, msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.log(ERROR, msg, args...)
}

func (l *Logger) PrintToFile(msg string) {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println(msg)
	log.SetOutput(os.Stdout) // Reset output to stdout
}
