package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/capstone-project-bunker/backend/services/auth/cmd/db"
	userDB "github.com/capstone-project-bunker/backend/services/auth/cmd/db/users"
	"github.com/capstone-project-bunker/backend/services/auth/internal/handlers"
	"github.com/capstone-project-bunker/backend/services/auth/pkg/constants/envKeys"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var authHandler *handlers.AuthHandler
var postgresUserDB *sql.DB
var userDBQueries *userDB.Queries

func init(){
	var err error
	if os.Getenv(envKeys.APP_ENV) == envKeys.PRODUCTION{
		err = godotenv.Load("prod.env")
	} else {
		err = godotenv.Load("dev.env")
	}
	if err != nil {
		log.Fatal(err)
	}
	
	postgresUserDB = db.GetDatabase()
	userDBQueries = userDB.New(postgresUserDB)
	authHandler = handlers.NewAuthHandler(userDBQueries)
}

func main(){
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.Writer.WriteString("yoyoyo")
	})
	r.POST("/login", authHandler.Login)
	r.POST("/validate", authHandler.Validate)
	
	defer postgresUserDB.Close()
	log.Fatal(r.Run(":" + os.Getenv("DEV_PORT")))
}