package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/capstone-project-bunker/backend/services/users/cmd/cache"
	dbPackage "github.com/capstone-project-bunker/backend/services/users/cmd/db"
	userQ "github.com/capstone-project-bunker/backend/services/users/cmd/db/queries/user"
	"github.com/capstone-project-bunker/backend/services/users/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var db *sql.DB
var userHandler *handlers.UserHandler
var rdb *redis.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db = dbPackage.GetDatabase()
	queriesDB := userQ.New(db)
	rdb = cache.NewRedisClient()
	
	userHandler = handlers.NewUserHandler(queriesDB, rdb)
}

func main() {
	r := gin.Default()
	r.POST("/", userHandler.Register)
	r.GET("/", userHandler.GetAll)
	r.GET("/by-email", userHandler.GetByEmail)
	r.DELETE("/by-email", userHandler.DeleteByEmail)
	r.DELETE("/:id", userHandler.DeleteById)

	defer db.Close()
	defer rdb.Close()
	log.Fatal(r.Run(":" + os.Getenv("DEV_PORT")))

}
