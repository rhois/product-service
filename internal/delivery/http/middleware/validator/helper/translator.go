package validatorhelper

import (
	"net/http"

	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

// MessageError DataStructure MessageError with field and message
type MessageError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

//CodeType define code error in every error
type CodeType string

var (
	// InvalidArgument type for the problem on input or validation
	InvalidArgument = CodeType("invalid-argument")
	// FatalError type for the problem on function (BUG) or other service
	FatalError = CodeType("fatal-error")
	// NotAuthenticate type for user data not authenticate
	NotAuthenticate = CodeType("not-authenticate")
	// BadRequest - when server will not process request because it's obvious client fault
	BadRequest = CodeType("bad-request")
)

// ResponseError response error struct to support format for handler
type ResponseError struct {
	OK              bool           `json:"ok"`
	Code            CodeType       `json:"code"`
	MessageResponse string         `json:"message"`
	ErrorMessages   []MessageError `json:"errors"`
}

// ErrorMessageTranslator function to customize the error message
func ErrorMessageTranslator(err error, translated ut.Translator) []MessageError {
	var messages []MessageError
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		var messageStruct MessageError
		messageStruct.Field = e.Field()
		messageStruct.Message = e.Translate(translated)
		messages = append(messages, messageStruct)
		// can translated each error one at a time.

	}
	return messages
}

// GetHTTPCode function for get httpcode from internal code type
func GetHTTPCode(Code CodeType) int {
	var httpCode int
	switch Code {
	case BadRequest:
		httpCode = http.StatusBadRequest
		break
	case FatalError:
		httpCode = http.StatusInternalServerError
		break
	case InvalidArgument:
		httpCode = http.StatusUnprocessableEntity
		break
	case NotAuthenticate:
		httpCode = http.StatusUnauthorized
		break
	}
	return httpCode
}
