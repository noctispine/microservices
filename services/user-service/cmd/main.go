package main

import (
	"database/sql"
	"log"
	"os"

	dbPackage "github.com/capstone-project-bunker/backend/cmd/db"
	userQ "github.com/capstone-project-bunker/backend/cmd/db/queries/user"
	"github.com/capstone-project-bunker/backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var db *sql.DB
var userHandler *handlers.UserHandler

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db = dbPackage.GetDatabase()
	queriesDB := userQ.New(db)
	userHandler = handlers.NewUserHandler(queriesDB)
}

func main() {
	r := gin.Default()
	r.POST("/", userHandler.Create)
	r.GET("/", userHandler.GetAll)
	r.GET("/by-email", userHandler.GetByEmail)

	log.Fatal(r.Run(":" + os.Getenv("DEV_PORT")))
}