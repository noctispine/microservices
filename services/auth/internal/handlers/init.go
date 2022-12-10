package handlers

import (
	"github.com/capstone-project-bunker/backend/services/auth/internal/validatorTranslations"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func init(){
	validatorTranslations.Validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	validatorTranslations.EnTrans, _ = uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validatorTranslations.Validate, validatorTranslations.EnTrans)
}