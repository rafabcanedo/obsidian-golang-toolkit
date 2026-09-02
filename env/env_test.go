package env_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/env"
)

func TestGetString(t *testing.T) {
	tests := []struct {
		name string
		key string
		envValue string
		defaultValue string
		want string
	} {
		{
			name:         "returns environment value when it is set",
			key:          "OBSIDIAN_TEST_STRING_VALUE",
			envValue:     "production",
			defaultValue: "development",
			want:         "production",
		},
		{
			name:         "returns default value when environment value is empty",
			key:          "OBSIDIAN_TEST_STRING_EMPTY",
			envValue:     "",
			defaultValue: "development",
			want:         "development",
		},
		{
			name:         "returns default value when environment variable is not set",
			key:          "OBSIDIAN_TEST_STRING_UNSET",
			envValue:     "",
			defaultValue: "development",
			want:         "development",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "returns default value when environment variable is not set" {
				t.Setenv(tt.key, "")
			} else {
				t.Setenv(tt.key, tt.envValue)
			}

			got := env.GetString(tt.key, tt.defaultValue)

			if got != tt.want {
				t.Errorf("GetString(%q, %q) = %q; want %q",
				    tt.key,
				    tt.defaultValue,
					got,
					tt.want,
			)
			}
		})
	}
}