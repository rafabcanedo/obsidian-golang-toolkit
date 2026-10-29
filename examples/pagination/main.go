package main

import (
	"log"
	"net/http"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/respond"
	"github.com/rafabcanedo/obsidian-golang-toolkit/pagination"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Rafael"},
	{ID: 2, Name: "Ana"},
	{ID: 3, Name: "Carlos"},
	{ID: 4, Name: "Maria"},
	{ID: 5, Name: "João"},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		p := pagination.New(1, 2)

		start := p.Offset()
		end := start + p.PageSize
		if end > len(users) {
			end = len(users)
		}

		meta := p.NewMeta(len(users))
		respond.Paginated(w, users[start:end], meta)
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
