package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/binding"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ListParams struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"search"`
}

func main() {
	http.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserRequest
		if err := binding.JSON(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "created: %s <%s>", req.Name, req.Email)
	})

	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		var params ListParams
		if err := binding.Query(r, &params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "page=%d limit=%d search=%q", params.Page, params.Limit, params.Search)
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
