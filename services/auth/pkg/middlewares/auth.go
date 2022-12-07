package middlewares

import (
	"net/http"
	"os"

	"github.com/capstone-project-bunker/backend/services/auth/internal/handlers"
	"github.com/capstone-project-bunker/backend/services/auth/pkg/constants/keys"
	"github.com/capstone-project-bunker/backend/services/auth/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)



func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := &handlers.Claims{}
		
		tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if tkn == nil || !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(keys.UserID, claims.UserID)
		c.Set(keys.UserRole, claims.Role)
		c.Next()
	}
}

func Authorization(avaliableRoles []int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.Contains(avaliableRoles, c.GetInt(keys.UserRole)){
			c.Next()
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}