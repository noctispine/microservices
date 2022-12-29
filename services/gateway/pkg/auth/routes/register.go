package routes

import (
	"log"
	"net/http"

	"github.com/capstone-project-bunker/backend/services/gateway/internal/validatorTranslations"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/auth/pb"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/responses"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required,max=50"`
	Surname  string `json:"surname" validate:"required,max=50"`
}

func Register(c *gin.Context, client pb.AuthServiceClient) {
	registerRequestBody := RegisterRequestBody{}

	if err := c.ShouldBindJSON(&registerRequestBody); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		// log.Error(err)
		return
	}

	if err := validatorTranslations.Validate.Struct(registerRequestBody) ; err != nil {
		responses.AbortWithStatusJSONValidationErrors(c, http.StatusBadRequest, err)
		return
	}

	res, err := client.Register(c, &pb.RegisterRequest{
		Email: registerRequestBody.Email,
		Password: registerRequestBody.Password,
		Name: registerRequestBody.Name,
		Surname: registerRequestBody.Surname,
	})

	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res.BaseResponse.Error != "" {
		responses.AbortWithStatusJSONError(c, res.BaseResponse.Status, err)
		return
	}

	c.Status(int(res.BaseResponse.Status))	
}