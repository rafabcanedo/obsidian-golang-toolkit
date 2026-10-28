package validator_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/validator"
)

func TestRequired(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"empty string", "", true},
		{"only spaces", "   ", true},
		{"valid value", "John", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.New().Required("field", tt.value).Validate()
			testutil.Equal(t, errs != nil, tt.want)
		})
	}
}

func TestMinLen(t *testing.T) {
	tests := []struct {
		name  string
		value string
		min   int
		want  bool
	}{
		{"too short", "ab", 3, true},
		{"exact length", "abc", 3, false},
		{"longer", "abcd", 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.New().MinLen("field", tt.value, tt.min).Validate()
			testutil.Equal(t, errs != nil, tt.want)
		})
	}
}

func TestMaxLen(t *testing.T) {
	tests := []struct {
		name  string
		value string
		max   int
		want  bool
	}{
		{"too long", "abcd", 3, true},
		{"exact length", "abc", 3, false},
		{"shorter", "ab", 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.New().MaxLen("field", tt.value, tt.max).Validate()
			testutil.Equal(t, errs != nil, tt.want)
		})
	}
}

func TestEmail(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"valid email", "user@example.com", false},
		{"missing @", "userexample.com", true},
		{"missing domain", "user@", true},
		{"empty", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.New().Email("email", tt.value).Validate()
			testutil.Equal(t, errs != nil, tt.want)
		})
	}
}

func TestMultipleErrors(t *testing.T) {
	errs := validator.New().
		Required("name", "").
		Email("email", "invalid").
		Validate()

	testutil.Equal(t, len(errs), 2)
}

func TestNoErrors(t *testing.T) {
	errs := validator.New().
		Required("name", "John").
		Email("email", "john@example.com").
		Validate()

	testutil.Equal(t, errs == nil, true)
}
