package config_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/config"
)

func TestLoadEnv(t *testing.T) {
	result, err := config.LoadEnv("testdata/sample.env")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["APP_ENV"] != "production" {
		t.Errorf("APP_ENV = %q; want %q", result["APP_ENV"], "production")
	}
	if result["APP_PORT"] != "8080" {
		t.Errorf("APP_PORT = %q; want %q", result["APP_PORT"], "8080")
	}
}

func TestLoadJSON(t *testing.T) {
	var result map[string]string
	err := config.LoadJSON("testdata/sample.json", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["app_env"] != "production" {
		t.Errorf("app_env = %q; want %q", result["app_env"], "production")
	}
}

func TestLoadYAML(t *testing.T) {
	var result map[string]string
	err := config.LoadYAML("testdata/sample.yaml", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["app_env"] != "production" {
		t.Errorf("app_env = %q; want %q", result["app_env"], "production")
	}
}
