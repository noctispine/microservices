package main

import (
	"shared/cutils"
	"shared/clogger"
	"github.com/capstone-project-bunker/backend/services/$serviceName/internal/config"
)

func init(){
	cutils.Logger = clogger.New("$serviceName", os.Getenv("APP_ENV"))
    cutils.Logger.Sugar().Infof("init success with env %s", "formatting works")

    var err error
    config.Conf, err = config.LoadConfig()
    if err != nil {
        utils.Logger.Sugar().Fatal("failed at config %w", err)
    }
}

func main() {

	listen, err := net.Listen("tcp", ":" + config.Conf.PORT)

	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	s := handlers.New$serviceNameCapitalized(userDBQueries)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor))

	pb.RegisterAuthServiceServer(grpcServer, s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
	
	defer postgresUserDB.Close()
    
    r := gin.Default()
    r.Use(middlewares.CORSMiddleware(), middlewares.AddCorelationID())

    authSvc := auth.RegisterRoutes(r, &config.Conf)
    users.RegisterRoutes(r, &config.Conf, authSvc)

    r.Run(config.Conf.Port)
}