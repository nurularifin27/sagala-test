package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the main logger instance.
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON for structured logging
}
