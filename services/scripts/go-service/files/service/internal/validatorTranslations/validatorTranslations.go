package validatorTranslations

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate
var EnTrans ut.Translator

func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
	  return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
	  translatedErr := fmt.Errorf(e.Translate(trans))
	  errs = append(errs, translatedErr)
	}
	return errs
}

func StringfyJSONErrArr(errs []error) []string {
	strErrors := make([]string, len(errs))

	for i, err := range errs {
		strErrors[i] = err.Error()
	}

	return strErrors
}