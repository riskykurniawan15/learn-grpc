package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()
	
	// Register custom validators
	v.RegisterValidation("alpha_space", validateAlphaSpace)
	v.RegisterValidation("password_strength", validatePasswordStrength)
	v.RegisterValidation("unique_email", validateUniqueEmail)
	
	return &Validator{
		validate: v,
	}
}

// ValidateStruct validates a struct
func (v *Validator) ValidateStruct(s interface{}) error {
	return v.validate.Struct(s)
}

// ValidateVar validates a single field
func (v *Validator) ValidateVar(field interface{}, tag string) error {
	return v.validate.Var(field, tag)
}

// Custom validation: alpha + space only
func validateAlphaSpace(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	matched, _ := regexp.MatchString(`^[a-zA-Z\s]+$`, value)
	return matched
}

// Custom validation: password strength
func validatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	
	// At least 8 characters
	if len(password) < 8 {
		return false
	}
	
	// At least one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	
	// At least one lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	
	// At least one number
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	
	// At least one special character
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)
	
	return hasUpper && hasLower && hasNumber && hasSpecial
}

// Custom validation: unique email (requires database check)
func validateUniqueEmail(fl validator.FieldLevel) bool {
	// This would need database context
	// For now, return true and handle uniqueness in service layer
	return true
}

// GetValidationErrors returns formatted validation errors
func (v *Validator) GetValidationErrors(err error) []string {
	var errors []string
	
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, formatValidationError(e))
		}
	}
	
	return errors
}

// Format validation error messages
func formatValidationError(e validator.FieldError) string {
	field := strings.ToLower(e.Field())
	
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, e.Param())
	case "alpha_space":
		return fmt.Sprintf("%s can only contain letters and spaces", field)
	case "password_strength":
		return fmt.Sprintf("%s must contain at least 8 characters with uppercase, lowercase, number, and special character", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
