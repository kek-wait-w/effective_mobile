package logs

import (
	"github.com/sirupsen/logrus"
)

var Logger = LoggerInit()

func LoggerInit() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	return logger
}

func LogError(logger *logrus.Logger, packageName, functionName string, err error, message string) {
	logger.WithFields(logrus.Fields{
		"package":  packageName,
		"function": functionName,
		"error":    err,
	}).Error(message)
}

func LogFatal(logger *logrus.Logger, packageName, functionName string, err error, message string) {
	logger.WithFields(logrus.Fields{
		"package":  packageName,
		"function": functionName,
		"error":    err,
	}).Fatal(message)
}
