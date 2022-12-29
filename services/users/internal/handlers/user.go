package handlers

import (
	"context"
	"net/http"

	userDB "github.com/capstone-project-bunker/backend/services/users/cmd/db/queries/user"
	"github.com/capstone-project-bunker/backend/services/users/pkg/pb"
	"github.com/capstone-project-bunker/backend/services/users/pkg/utils"
	"github.com/capstone-project-bunker/backend/services/users/pkg/wrappers"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	db  *userDB.Queries
	rdb *redis.Client
}

func NewUserService(db *userDB.Queries, rdb *redis.Client) *UserService {
	return &UserService{
		db,
		rdb,
	}
}

func (h *UserService) GetAll(c context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error){
	var dbUsers []userDB.User
	var err error

	dbUsers, err = h.db.GetAll(c)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var res pb.GetAllResponse

	for _, user := range dbUsers {
		res.Users = append(res.Users, &pb.User{Id: user.ID.String(),
			Email: user.Email,
			Name: user.Name,
			Surname: user.Surname,
			Role: pb.ROLES(user.Role),
			CreatedAt: timestamppb.New(user.CreatedAt),
			LastLoginAt: timestamppb.New(user.LastLoginAt.Time),
			IsActive: user.IsActive,})
	}

	res.BaseResponse = &pb.Response{
		Status: http.StatusOK,
	}

	return &res, nil
}

func (h *UserService) GetById(c context.Context, req *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.GetByIdResponse{
			BaseResponse: &pb.Response{
				Status: http.StatusBadRequest,
				Error: wrappers.NewErrNotValid("user id").Error(),
			},
		}, nil
	}

	user, err := h.db.GetById(c, id)

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.GetByIdResponse{
		BaseResponse: &pb.Response{
			Status: http.StatusOK,
		},
		User: &pb.User{
			Id: user.ID.String(),
			Email: user.Email,
			Name: user.Name,
			Surname: user.Surname,
			Role: pb.ROLES(user.Role),
			CreatedAt: timestamppb.New(user.CreatedAt),
			LastLoginAt: timestamppb.New(user.LastLoginAt.Time),
			IsActive: user.IsActive,
		},
	}, nil
}

func (h *UserService) DeleteById(c context.Context, req *pb.DeleteByIdRequest) (*pb.DeleteByIdResponse, error) {
	// idString := c.Params.ByName("id")
	// if idString == "" {
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.DeleteByIdResponse{
			BaseResponse: &pb.Response{
				Status: http.StatusBadRequest,
				Error: wrappers.NewErrNotValid("user id").Error(),
			},
		}, nil
	}

	rowsAffected, err := h.db.DeleteById(c, id)
	if rowsAffected == 0 {
		return &pb.DeleteByIdResponse{
			BaseResponse: &pb.Response{
				Status: http.StatusNotFound,
				Error: wrappers.NewErrNotFound("user").Error(),
			},
		}, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.DeleteByIdResponse{
		BaseResponse: &pb.Response{
			Status: http.StatusNoContent,
		},
	}, nil
}

func (h *UserService) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error){
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error()) 
	}

	createUserParams := userDB.CreateParams{
		Email: req.Email,
		HashedPassword: hashedPassword,
		Name: req.Name,
		Surname: req.Surname,
	}

	if err := h.db.Create(c, createUserParams); err != nil {
		if utils.CheckPostgreError(err, pgerrcode.UniqueViolation) {
			return &pb.RegisterResponse{
				BaseResponse: &pb.Response{
					Status: http.StatusBadRequest,
					Error: wrappers.NewErrAlreadyExists("user").Error(),
				},
			}, nil
		}

		return nil, status.Error(codes.Unknown, err.Error())
	}
	
	return &pb.RegisterResponse{
		BaseResponse: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
} 

func (h *UserService) ActivateUser(c context.Context, req *pb.ActivateUserRequest) (*pb.ActivateUserResponse, error) {
	// idString := c.Params.ByName("id")
	// if idString == "" {
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }

	id, err := uuid.Parse(req.Id)

	if err != nil {
		return &pb.ActivateUserResponse{
			BaseResponse: &pb.Response{
				Status: http.StatusBadRequest,
				Error: wrappers.NewErrNotValid("user id").Error(),
			},
		}, nil
	}

	rowsAffected, err := h.db.ActivateUser(c, id)

	if rowsAffected == 0 {
		return &pb.ActivateUserResponse{
			BaseResponse: &pb.Response{
				Status: http.StatusNotFound,
				Error: wrappers.NewErrNotFound("user").Error(),
			},
		}, nil
	}
	
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	
	return &pb.ActivateUserResponse{
		BaseResponse: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}
