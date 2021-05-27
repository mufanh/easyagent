package global

import (
	"github.com/mufanh/easyagent/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"path/filepath"
)

var (
	Logger *logger.Logger
)

func SetupLogger(logFilepath string, logFilename string, maxSize int, maxAge int) error {
	Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  filepath.Join(logFilepath, logFilename),
			MaxSize:   maxSize,
			MaxAge:    maxAge,
			LocalTime: true,
		},
		"",
		log.LstdFlags).WithCaller(2)
	return nil
}
