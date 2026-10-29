package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/rafabcanedo/obsidian-golang-toolkit/health"
)

func main() {
	h := health.New().
		AddCheck("database", func(ctx context.Context) error {
			return nil
		}).
		AddCheck("cache", func(ctx context.Context) error {
			return errors.New("cache unreachable")
		})

	http.Handle("/health", h.Handler())

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
