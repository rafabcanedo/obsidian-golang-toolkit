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

func TestMin(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		min   float64
		want  bool
	}{
		{"below min", -1, 0, true},
		{"at min", 0, 0, false},
		{"above min", 1, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.New().Min("field", tt.value, tt.min).Validate()
			testutil.Equal(t, errs != nil, tt.want)
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		max   float64
		want  bool
	}{
		{"above max", 32, 31, true},
		{"at max", 31, 31, false},
		{"below max", 1, 31, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.New().Max("field", tt.value, tt.max).Validate()
			testutil.Equal(t, errs != nil, tt.want)
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		min   float64
		max   float64
		want  bool
	}{
		{"below range", 0, 1, 31, true},
		{"above range", 32, 1, 31, true},
		{"at min", 1, 1, 31, false},
		{"at max", 31, 1, 31, false},
		{"in range", 15, 1, 31, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.New().Between("field", tt.value, tt.min, tt.max).Validate()
			testutil.Equal(t, errs != nil, tt.want)
		})
	}
}

func TestNoErrors(t *testing.T) {
	errs := validator.New().
		Required("name", "John").
		Email("email", "john@example.com").
		Validate()

	testutil.Equal(t, errs == nil, true)
}
