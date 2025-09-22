package middleware

import (
	"net/http"
	"sagala/pkg/utils"
	"sagala/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if err == gorm.ErrRecordNotFound {
				utils.NotFound(c, "Resource not found")
				c.Abort()
				return
			}

			switch e := err.(type) {
			case validator.ValidationErrors:
				errors := validation.FormatValidationErrors(e)
				utils.ErrorWithDetails(c, http.StatusBadRequest, "Validation failed", errors)
				c.Abort()
				return

			default:
				utils.Error(c, http.StatusInternalServerError, "An unexpected error occurred")
				c.Abort()
				return
			}
		}
	}
}
