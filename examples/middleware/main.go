package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := middleware.GetRequestID(r)
		fmt.Fprintf(w, "request_id=%s", id)
	})

	cors := middleware.CORS(middleware.CORSOptions{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := middleware.Logger(
		middleware.Recovery(
			middleware.RequestID(
				cors(mux),
			),
		),
	)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
