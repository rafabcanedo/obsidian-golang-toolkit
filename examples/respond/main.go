package main

import (
	"log"
	"net/http"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/respond"
	"github.com/rafabcanedo/obsidian-golang-toolkit/pagination"
	"github.com/rafabcanedo/obsidian-golang-toolkit/util/errutil"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		respond.OK(w, User{ID: 1, Name: "Rafael"})
	})

	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{{ID: 1, Name: "Rafael"}, {ID: 2, Name: "Ana"}}
		meta := pagination.New(1, 10).NewMeta(2)
		respond.Paginated(w, users, meta)
	})

	http.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		respond.Created(w, User{ID: 3, Name: "Carlos"})
	})

	http.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		respond.NoContent(w)
	})

	http.HandleFunc("GET /users/missing", func(w http.ResponseWriter, r *http.Request) {
		respond.Error(w, errutil.NotFound("user not found"))
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
