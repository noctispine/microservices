package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"strconv"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/capstone-project-bunker/backend/services/$serviceName/pkg/pb"
)

type $serviceNameCapitalizedService struct {
	pb.Unimplemented$serviceNameCapitalizedServiceServer
}
func New$serviceNameCapitalizedService() *$serviceNameCapitalizedService {
	return &$serviceNameCapitalizedService{}

}

func (h *$serviceNameCapitalizedService) Get(c context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return  nil, status.Error(codes.Unknown, err.Error())
}