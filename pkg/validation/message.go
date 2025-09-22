package validation

var validationMessageMap = map[string]string{
	"required": "This field is required.",
	"email":    "Invalid email format.",
	"min":      "The value must be at least %s.",
	"max":      "The value must not exceed %s.",
	"len":      "The value must be exactly %s characters long.",
	"gt":       "The value must be greater than %s.",
	"gte":      "The value must be greater than or equal to %s.",
	"lt":       "The value must be less than %s.",
	"lte":      "The value must be less than or equal to %s.",
}
