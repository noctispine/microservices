package clogger

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

func New(service string, env string) *zap.Logger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger.With(zap.String("service", service)).With(zap.String("environment", env))
}
