package main

import (
	"log"

	"github.com/capstone-project-bunker/backend/services/gateway/pkg/auth"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/config"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/users"
	"github.com/gin-gonic/gin"
)



func main() {
    config, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

    r := gin.Default()

    authSvc := auth.RegisterRoutes(r, &config)
    users.RegisterRoutes(r, &config, authSvc)

    r.Run(config.Port)
}