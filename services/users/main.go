package main

import (
	"log"
	"net"
	"os"

	"github.com/capstone-project-bunker/backend/services/users/cmd/cache"
	dbPackage "github.com/capstone-project-bunker/backend/services/users/cmd/db"
	userDB "github.com/capstone-project-bunker/backend/services/users/cmd/db/queries/user"
	"github.com/capstone-project-bunker/backend/services/users/internal/handlers"
	"github.com/capstone-project-bunker/backend/services/users/pkg/constants/envKeys"
	"github.com/capstone-project-bunker/backend/services/users/pkg/pb"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var rdb *redis.Client
var queriesDB *userDB.Queries

func init() {
	var err error
	
	if os.Getenv(envKeys.APP_ENV) == envKeys.PRODUCTION {
		err = godotenv.Load("prod.env")
	} else {
		err = godotenv.Load("dev.env")
	}

	if err != nil {
		log.Fatal(err)
	}

	db := dbPackage.GetDatabase()
	queriesDB = userDB.New(db)
	rdb = cache.NewRedisClient()
}

func main() {
	listen, err := net.Listen("tcp", ":" + os.Getenv("PORT"))

	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	service := handlers.NewUserService(queriesDB, rdb)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, service)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
	
	// r.POST("/", userHandler.Register)
	// r.GET("/", userHandler.GetAll)
	// r.GET("/by-email", userHandler.GetByEmail)
	// r.DELETE("/by-email", userHandler.DeleteByEmail)
	// r.DELETE("/:id", userHandler.DeleteById)
	// r.PATCH("/:id", userHandler.ActivateUser)

	// defer db.Close()
	defer rdb.Close()
}
