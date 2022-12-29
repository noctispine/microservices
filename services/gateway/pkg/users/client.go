package users

import (
	"fmt"

	"github.com/capstone-project-bunker/backend/services/gateway/pkg/config"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/users/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}

func InitServiceClient(c *config.Config) pb.UserServiceClient {
	var cc *grpc.ClientConn
	var err error

	// if utils.CheckIsDevelopment() {
	// 	cc, err = grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// } else {
	// 	cc, err = grpc.Dial(c.AuthSvcUrl)
	// }

	cc, err = grpc.Dial(c.UsersSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewUserServiceClient(cc)

}