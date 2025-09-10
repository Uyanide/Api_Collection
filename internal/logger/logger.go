package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// InitLogger initializes the global logger
func InitLogger() {
	Log = logrus.New()

	Log.SetOutput(os.Stdout)

	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Log.SetLevel(logrus.InfoLevel)
}

func GetLogger() *logrus.Logger {
	if Log == nil {
		InitLogger()
	}
	return Log
}
