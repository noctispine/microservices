package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/capstone-project-bunker/backend/services/users/cmd/db"
	userDB "github.com/capstone-project-bunker/backend/services/users/cmd/db/queries/user"
	"github.com/capstone-project-bunker/backend/services/users/pkg/constants/envKeys"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

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
	// authHandler = handlers.NewAuthHandler(userDBQueries)
}

func main(){
	listen, err := net.Listen("tcp", ":" + os.Getenv("PORT"))

	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	s := handlers.NewAuthService(userDBQueries)

	grpcServer := grpc.NewServer()

	authPB.RegisterAuthServiceServer(grpcServer, s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
	
	defer postgresUserDB.Close()
}