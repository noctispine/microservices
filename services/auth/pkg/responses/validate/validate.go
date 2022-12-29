package validateResp

import (
	"github.com/capstone-project-bunker/backend/services/auth/pkg/pb"
)

func ConstructBaseRespWithErr(status int32, err error) *pb.ValidateResponse {
	return &pb.ValidateResponse{
		BaseResponse: &pb.Response{
			Status: status,
			Error: err.Error(),
		},
	}
}

func ConstructBaseResp(status int32) *pb.ValidateResponse {
	return &pb.ValidateResponse{
		BaseResponse: &pb.Response{
			Status: status,
		},
	}
}

