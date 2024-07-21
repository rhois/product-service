package httpvalidator

import (
	"net/http"
	"product/internal/presenter"

	validatorHelper "product/internal/delivery/http/middleware/validator/helper"

	"github.com/gin-gonic/gin"
	v "gopkg.in/go-playground/validator.v9"
)

// ValidateCreateProduct is function to validate product data from payload
func ValidateCreateProduct(c *gin.Context) {
	var postData presenter.Products
	c.BindJSON(&postData)
	validate := v.New()
	translated := validatorHelper.RequiredErrorMessage(validate)
	err := validate.Struct(postData)
	if err != nil {
		errorMessages := validatorHelper.ErrorMessageTranslator(err, translated)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"code":    "validation-error",
			"message": "Required field(s) is not supplied or having an invalid type please check again",
			"errors":  errorMessages,
		})
		return
	}
	c.Set("post_data", postData)
	c.Next()
}
