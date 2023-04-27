package main

import (
	"os"

	"shared/clogger"

	"github.com/capstone-project-bunker/backend/services/gateway/pkg/auth"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/config"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/middlewares"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/users"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)


func init() {
    utils.Logger = clogger.New("gateway", os.Getenv("APP_ENV"))
    utils.Logger.Sugar().Infof("init success with env %s", "formatting works")

    var err error
    config.Conf, err = config.LoadConfig()
    if err != nil {
        utils.Logger.Sugar().Fatal("failed at config %w", err)
    }
}

func main() {
    
    r := gin.Default()
    r.Use(middlewares.CORSMiddleware(), middlewares.AddCorelationID())

    authSvc := auth.RegisterRoutes(r, &config.Conf)
    users.RegisterRoutes(r, &config.Conf, authSvc)

    r.Run(config.Conf.Port)
}