package logger

import (
	"os"

	"github.com/rs/zerolog"
)

const (
	logDir   = "logs/"
	fileMode = 0666
)

func Init(level string) zerolog.Logger {
	path := logDir + "log.txt"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, fileMode)

	if err != nil {
		panic(err)
	}

	var logLevel zerolog.Level
	if level == "debug" {
		logLevel = zerolog.DebugLevel
	} else {
		logLevel = zerolog.WarnLevel
	}

	logger := zerolog.New(file).Level(logLevel)
	return logger
}
