package auth

import (
	"fmt"
	"net/http"

	"github.com/capstone-project-bunker/backend/services/gateway/pkg/auth/pb"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/constants/keys"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (config *AuthMiddlewareConfig) ValidateToken(c *gin.Context) {
	tkn := c.GetHeader("Authorization")

	if tkn == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}	

	res, err := config.svc.Client.Validate(c, &pb.ValidateRequest{
		Token: tkn,
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Println(res.Id, res.Role)
	c.Set(keys.UserID, res.Id)
	c.Set(keys.UserRole, res.Role)
	c.Next()
}

