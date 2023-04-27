package cutils

import (
	"os"
	"shared/cconstants"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func CheckIsProd() bool {
	return os.Getenv(cconstants.ENV.APP_ENV) == cconstants.ENV.PRODUCTION
}