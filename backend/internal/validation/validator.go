package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct and returns formatted errors
func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// FormatValidationErrors formats validation errors for API response
func FormatValidationErrors(err error) []ValidationError {
	var errors []ValidationError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   e.Field(),
				Message: getErrorMessage(e),
			})
		}
	}

	return errors
}

// getErrorMessage returns a user-friendly error message
func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short (minimum " + e.Param() + " characters)"
	case "max":
		return "Value is too long (maximum " + e.Param() + " characters)"
	case "eqfield":
		return "Value must match " + e.Param()
	default:
		return "Invalid value"
	}
}

// SendValidationError sends a validation error response
func SendValidationError(c *fiber.Ctx, err error) error {
	errors := FormatValidationErrors(err)
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    "VALIDATION_ERROR",
			"message": "Validation failed",
			"details": errors,
		},
	})
}
