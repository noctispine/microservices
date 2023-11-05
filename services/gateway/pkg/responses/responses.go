package responses

import (
	"github.com/capstone-project-bunker/backend/services/gateway/internal/validatorTranslations"
	"github.com/gin-gonic/gin"
)



func AbortWithStatusJSONError(c *gin.Context, code int32, err error) {
	c.AbortWithStatusJSON(int(code), gin.H{
		"error": err.Error()})
}

func AbortWithStatusJSONErrorMessage(c *gin.Context, code int32, errMessage string) {
	c.AbortWithStatusJSON(int(code), gin.H{
		"error": errMessage})
}

func AbortWithStatusJSONValidationErrors(c *gin.Context, code int, err error) {
	errs := validatorTranslations.TranslateError(err, validatorTranslations.EnTrans)
    c.AbortWithStatusJSON(code , gin.H{
		"errors": validatorTranslations.StringfyJSONErrArr(errs)})
}