package strutil_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/util/strutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty string", "", true},
		{"only spaces", "   ", true},
		{"non-empty", "hello", false},
		{"space padded", " hi ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, strutil.IsEmpty(tt.input), tt.want)
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"lowercase", "hello", "Hello"},
		{"already capitalized", "Hello", "Hello"},
		{"empty", "", ""},
		{"single char", "a", "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, strutil.Capitalize(tt.input), tt.want)
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name  string
		input string
		max   int
		want  string
	}{
		{"no truncation needed", "Hello", 10, "Hello"},
		{"exact length", "Hello", 5, "Hello"},
		{"truncated", "Hello World", 5, "Hello..."},
		{"empty", "", 5, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, strutil.Truncate(tt.input, tt.max), tt.want)
		})
	}
}

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple camel", "UserName", "user_name"},
		{"multiple words", "MyFunctionName", "my_function_name"},
		{"already lower", "username", "username"},
		{"single word", "User", "user"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, strutil.CamelToSnake(tt.input), tt.want)
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", "Hello World", "hello-world"},
		{"with accents", "Olá Mundo", "ola-mundo"},
		{"with special chars", "Hello, World!", "hello-world"},
		{"multiple spaces", "Hello   World", "hello-world"},
		{"leading trailing spaces", "  hello  ", "hello"},
		{"with numbers", "Post 123", "post-123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, strutil.Slugify(tt.input), tt.want)
		})
	}
}
