package loginResp

import (
	"github.com/capstone-project-bunker/backend/services/auth/pkg/pb"
)

func ConstructBaseRespWithErr(status int, err error) *pb.LoginResponse {
	return &pb.LoginResponse{
		BaseResponse: &pb.Response{
			Status: int32(status),
			Error: err.Error(),
		},
	}
}

func ConstructBaseResp(status int) *pb.LoginResponse {
	return &pb.LoginResponse{
		BaseResponse: &pb.Response{
			Status: int32(status),
		},
	}
}

