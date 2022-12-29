package responses

import (
	"fmt"

	"github.com/capstone-project-bunker/backend/services/auth/internal/validatorTranslations"
	"github.com/gin-gonic/gin"
)

func AbortWithStatusJSONError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, gin.H{
		"error": err.Error()})
}

func AbortWithStatusJSONValidationErrors(c *gin.Context, code int, err error) {
	errs := validatorTranslations.TranslateError(err, validatorTranslations.EnTrans)
	fmt.Println(errs)
    c.AbortWithStatusJSON(code , gin.H{
		"errors": validatorTranslations.StringfyJSONErrArr(errs)})
}

