package handlers

import (
	"fmt"
	"log"
	"net/http"

	userQ "github.com/capstone-project-bunker/backend/cmd/db/queries/user"
	"github.com/capstone-project-bunker/backend/pkg/responses"
	"github.com/capstone-project-bunker/backend/pkg/utils"
	"github.com/capstone-project-bunker/backend/pkg/wrappers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
)

type UserHandler struct {
	db *userQ.Queries
	rdb *redis.Client
}

func NewUserHandler(db *userQ.Queries, rdb *redis.Client) *UserHandler {
	return &UserHandler{
		db,
		rdb,
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var user userQ.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}


	if err := validate.Struct(user); err != nil {
		responses.AbortWithStatusJSONValidationErrors(c, http.StatusBadRequest, err)
		return
	}

	if err := h.db.Create(c, userQ.CreateParams{
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	}); err != nil {

		if utils.CheckPostgreError(err, pgerrcode.UniqueViolation) {
			responses.AbortWithStatusJSONError(c, http.StatusBadRequest, wrappers.NewErrAlreadyExists("email"))
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	var users []userQ.User
	var err error

	users, err = h.db.GetAll(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) DeleteByEmail(c *gin.Context) {
	userEmailReq := struct {
		Email string `json:"email" db:"email" validate:"required,email"`
	}{}
	if err := c.ShouldBindJSON(&userEmailReq); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(userEmailReq); err != nil {
		responses.AbortWithStatusJSONError(c, http.StatusBadRequest, wrappers.NewErrNotValid("email"))
		return
	}

	rowsAffected, err := h.db.DeleteByEmail(c, userEmailReq.Email)
	if rowsAffected == 0 {
		responses.AbortWithStatusJSONError(c, http.StatusNotFound, wrappers.NewErrNotFound("user"))
		return
	}
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) GetByEmail(c *gin.Context) {
	userEmailReq := struct {
		Email string `json:"email" db:"email" validate:"required,email"`
	}{}

	if err := c.ShouldBindJSON(&userEmailReq); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(userEmailReq); err != nil {
		responses.AbortWithStatusJSONError(c, http.StatusBadRequest, wrappers.NewErrNotValid("email"))
		return
	}

	fmt.Println(userEmailReq)

	user, err := h.db.GetByEmail(c, userEmailReq.Email)
	if err != nil {
		if utils.CheckPostgreError(err, pgerrcode.NoDataFound) {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteById(c *gin.Context) {
	idString := c.Params.ByName("id")
	if idString == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}


	rowsAffected, err := h.db.DeleteById(c, id)
	if rowsAffected == 0 {
		responses.AbortWithStatusJSONError(c, http.StatusNotFound, wrappers.NewErrNotFound("user"))
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}