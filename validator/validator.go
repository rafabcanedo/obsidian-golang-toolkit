package validator

import (
	"fmt"
	"regexp"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

type Errors []ValidationError

func (e Errors) Error() string {
	msgs := make([]string, len(e))
	for i, err := range e {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, "; ")
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type Validator struct {
	errors Errors
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Required(field, value string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "is required"})
	}
	return v
}

func (v *Validator) MinLen(field, value string, min int) *Validator {
	if len([]rune(value)) < min {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be at least %d characters", min),
		})
	}
	return v
}

func (v *Validator) MaxLen(field, value string, max int) *Validator {
	if len([]rune(value)) > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be at most %d characters", max),
		})
	}
	return v
}

func (v *Validator) Email(field, value string) *Validator {
	if !emailRegex.MatchString(value) {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "must be a valid email"})
	}
	return v
}

func (v *Validator) Validate() Errors {
	if len(v.errors) == 0 {
		return nil
	}
	return v.errors
}
