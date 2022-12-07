package handlers

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/capstone-project-bunker/backend/services/auth/cmd/db/users"
	"github.com/capstone-project-bunker/backend/services/auth/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	Email string `json:"email"`
	UserID uuid.UUID 
	jwt.RegisteredClaims
	Role int32
}

type AuthHandler struct {
	ctx context.Context
	db *users.Queries
}



func (h *AuthHandler) SignInHandler(c *gin.Context) {
	user := struct {
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6,max=64"`
	}{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	
	// if err := validate.Struct(user); err != nil {
	// 	responses.AbortWithStatusJSONValidationErrors(c, http.StatusBadRequest, err)
	// 	return
	// }

	dbUser, err := h.db.GetByEmail(c, user.Email)
	if err != nil {
		// notfound implementation



		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}


	// if !dbUser.IsActive {
	// 	responses.AbortWithStatusJSONError(c, http.StatusBadRequest, fmt.Errorf("user is not activated yet, you can contact via email on home page"))
	// 	return
	// }

	if !utils.CheckPasswordHash(user.Password, dbUser.HashedPassword) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong Credentials"})
		return
	}

	expireInMinutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_MINUTES"))
	if err != nil {
		// log.Error(fmt.Errorf("jwt conversion to int: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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
		c.AbortWithStatus(http.StatusInternalServerError)
		// log.Error(fmt.Errorf("jwt signed string: %w", err))
		return
	}

	// dbUser.LastLoginAt = time.Now()
	// if err := h.db.Save(&dbUser).Error; err != nil {
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	// log.Error(err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"user": users.User{
			ID: dbUser.ID,
			Email: dbUser.Email,
			Role: dbUser.Role,
			Name: dbUser.Name,

		},
		"token": tokenString})
}