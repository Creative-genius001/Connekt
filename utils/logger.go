package utils

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger instance
var Logger = logrus.New()

func InitLogger() {
	// // Log output to file with rotation
	// Logger.SetOutput(&lumberjack.Logger{
	// 	Filename:   "logs/app.log", // Log file location
	// 	MaxSize:    10,             // Max size in MB before rotation
	// 	MaxBackups: 5,              // Max old log files to keep
	// 	MaxAge:     30,             // Max days to keep logs
	// 	Compress:   true,           // Gzip old logs
	// })

	// Set log format (JSON for structured logging)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// Log to both file and console
	Logger.SetOutput(os.Stdout)

	// Set log level (Debug, Info, Warn, Error)
	Logger.SetLevel(logrus.DebugLevel)
}

// Log functions
func Info(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Info(message)
}

func Warn(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Warn(message)
}

func Error(message string, err error, fields logrus.Fields) {
	Logger.WithFields(fields).WithError(err).Error(message)
}

func Debug(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Debug(message)
}
