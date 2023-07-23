package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func InitLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "[SocialVueGoApp] ", log.LstdFlags|log.Lshortfile)
	}
	return logger
}

func SetLogger(customLogger *log.Logger) {
	logger = customLogger
}

func LogError(format string, args ...interface{}) {
	if logger != nil {
		logger.Printf("[ERROR] "+format, args...)
	}
}

func LogInfo(format string, args ...interface{}) {
	if logger != nil {
		logger.Printf("[INFO] "+format, args...)
	}
}
