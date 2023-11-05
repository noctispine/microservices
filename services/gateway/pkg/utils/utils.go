package utils

import (
	"os"

	"github.com/capstone-project-bunker/backend/services/gateway/pkg/constants/envKeys"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func CheckIsProd() bool {
	return os.Getenv(envKeys.APP_ENV) == envKeys.PRODUCTION
}