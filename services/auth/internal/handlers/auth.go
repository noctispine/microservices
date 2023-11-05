package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"strconv"
	"time"

	userDB "github.com/capstone-project-bunker/backend/services/auth/cmd/db/users"
	"github.com/capstone-project-bunker/backend/services/auth/pkg/pb"
	"github.com/capstone-project-bunker/backend/services/auth/pkg/utils"
	"github.com/capstone-project-bunker/backend/services/auth/pkg/wrappers"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	db *userDB.Queries
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(db *userDB.Queries) *AuthService {
	return &AuthService{
		db: db,
	}
}
type Claims struct {
	Email string `json:"email"`
	UserID uuid.UUID 
	jwt.RegisteredClaims
	Role int32
}


// func response(resp *pb.LoginResponse,status int64, err error) (*pb.LoginResponse, error) {
// 	resp.BaseResponse.Error = err.Error()
// 	resp.BaseResponse.Status = status
// 	return resp, err
// }



func (h *AuthService) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	dbUser, err := h.db.GetByEmail(c, req.Email)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &pb.LoginResponse{
				BaseResponse: &pb.Response{
					Status: http.StatusBadRequest,
					Error: wrappers.NewErrDoesNotExist("user").Error(),
				}, 
			}, nil

		}

		return  nil, status.Error(codes.Unknown, err.Error())
	}

	// if !dbUser.IsActive {
	// 	return &pb.LoginResponse{
	// 		BaseResponse: &pb.Response{
	// 			Status: http.StatusBadRequest,
	// 			Error: fmt.Errorf("user is not activated yet").Error(),
	// 		},
	// 	}, status.Error(codes.FailedPrecondition, err.Error())
	// }

	if !utils.CheckPasswordHash(req.Password, dbUser.HashedPassword) {
		return &pb.LoginResponse{
			BaseResponse: &pb.Response{
				Status: http.StatusBadRequest,
				Error: "wrong credentials",
			},
		}, nil
	}

	expireInMinutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_MINUTES"))
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	
	expirationTime := time.Now().Add(time.Duration(expireInMinutes) * time.Minute)
	claims := &Claims{
		Email: dbUser.Email,
		UserID: dbUser.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		Role: dbUser.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
			return nil, status.Error(codes.Unknown, err.Error())
		}

		if err := h.db.UpdateLastLoginAt(c, dbUser.ID, time.Now()); err != nil {

			return nil, status.Error(codes.Unknown, err.Error())
		}

		return &pb.LoginResponse{
			Token: tokenString,
			Id: dbUser.ID.String(),
			Role: pb.ROLES(dbUser.Role),
			BaseResponse: &pb.Response{
				Status: http.StatusOK,
			},
		}, nil
	}

func (h *AuthService) Validate(c context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error){
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (interface{}, error){
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return &pb.ValidateResponse{
			BaseResponse: &pb.Response{
				Status: http.StatusUnauthorized,
				Error: wrappers.NewErrNotValid("token").Error(),
			},
		}, nil
	}

	if tkn == nil || !tkn.Valid {
		return &pb.ValidateResponse{
			BaseResponse: &pb.Response{
				Status: http.StatusUnauthorized,
				Error: wrappers.NewErrNotValid("token").Error(),
			},
		}, nil
	}
	
	return &pb.ValidateResponse{
		BaseResponse: &pb.Response{
			Status: http.StatusOK,	
		},
		Id: claims.ID,
		Role: pb.ROLES(claims.Role),
	}, nil
} 

func (h *AuthService) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error){
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
