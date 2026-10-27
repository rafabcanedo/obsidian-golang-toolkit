package logutil

import (
	"log/slog"
	"os"
)

func NewLogger(env string) *slog.Logger {
	if env == "production" {
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
