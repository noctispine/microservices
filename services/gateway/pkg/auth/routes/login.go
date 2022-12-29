package routes

import (
	"fmt"
	"net/http"

	"github.com/capstone-project-bunker/backend/services/gateway/internal/validatorTranslations"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/auth/pb"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/responses"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func Login(c *gin.Context, client pb.AuthServiceClient) {
	body := LoginRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := validatorTranslations.Validate.Struct(&body); err != nil {
		responses.AbortWithStatusJSONValidationErrors(c, http.StatusBadRequest, err)
		return
	}

	res, err := client.Login(c, &pb.LoginRequest{
		Email: body.Email,
		Password: body.Password,
	})

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if res.BaseResponse.Error != "" {
		responses.AbortWithStatusJSONErrorMessage(c, res.BaseResponse.Status, res.BaseResponse.Error)
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"token": res.Token,
		"id": res.Id,
		"role": res.Role,
	})
}
