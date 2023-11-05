package handlers

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var enTrans ut.Translator

func init(){
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	enTrans, _ = uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, enTrans)

}