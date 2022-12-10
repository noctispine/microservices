package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/capstone-project-bunker/backend/services/auth/cmd/db"
	userDB "github.com/capstone-project-bunker/backend/services/auth/cmd/db/users"
	"github.com/capstone-project-bunker/backend/services/auth/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var authHandler *handlers.AuthHandler
var postgresUserDB *sql.DB
var userDBQueries *userDB.Queries

func init(){
	err := godotenv.Load(".env")
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
	r.POST("/login", authHandler.SignInHandler)
	
	defer postgresUserDB.Close()
	log.Fatal(r.Run(":" + os.Getenv("DEV_PORT")))
}