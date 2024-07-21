package validatorhelper

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

// RequiredErrorMessage function to replace error message for  "required" type
func RequiredErrorMessage(validate *validator.Validate) ut.Translator {
	// NOTE: ommitting allot of error checking for brevity

	en := en.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	FieldJSONFormatter(validate)
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is Required!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.StructField())

		return t
	})
	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		errMsg := ""
		switch fe.Type().Name() {
		case "string":
			errMsg = "please fill with character length must be >= %s"
			break
		default:
			errMsg = "please fill with minimum value is %s"
			break
		}
		t, _ := ut.T("min", fmt.Sprintf(errMsg, fe.Param()))
		return t
	})
	validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		errMsg := ""
		switch fe.Type().Name() {
		case "string":
			errMsg = "please fill with character length must be <= %s"
			break
		default:
			errMsg = "please fill with maximum value is %s"
			break
		}
		t, _ := ut.T("max", fmt.Sprintf(errMsg, fe.Param()))
		return t
	})
	return trans
}
