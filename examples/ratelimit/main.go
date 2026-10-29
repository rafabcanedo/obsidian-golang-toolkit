package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/ratelimit"
)

func main() {
	limiter := ratelimit.New(10, time.Minute)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "here is your data")
	})

	log.Println("listening on :8080 (10 req/min per IP)")
	log.Fatal(http.ListenAndServe(":8080", limiter.Middleware()(mux)))
}
