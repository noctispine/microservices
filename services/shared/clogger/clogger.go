package clogger

import (
	"os"

	"go.uber.org/zap"
)
const logPath = "./logs/service.log"

func New(service string, env string) *zap.Logger {
	os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logPath}
	logger, _ := c.Build()
	return logger.With(zap.String("service", service)).With(zap.String("environment", env))
}
