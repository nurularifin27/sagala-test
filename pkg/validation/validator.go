package validation

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(errs validator.ValidationErrors) []gin.H {
	var errors []gin.H
	for _, fieldErr := range errs {
		message := fmt.Sprintf("Validation failed for field '%s'.", fieldErr.Field())

		if msg, ok := validationMessageMap[fieldErr.Tag()]; ok {
			if fieldErr.Param() != "" {
				message = fmt.Sprintf(msg, fieldErr.Param())
			} else {
				message = msg
			}
		}

		errors = append(errors, gin.H{
			"field":   fieldErr.Field(),
			"message": message,
		})
	}
	return errors
}
