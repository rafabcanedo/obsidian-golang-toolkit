package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/binding"
	"github.com/rafabcanedo/obsidian-golang-toolkit/api/respond"
	"github.com/rafabcanedo/obsidian-golang-toolkit/util/errutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/validator"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := binding.JSON(r, &req); err != nil {
			respond.Error(w, errutil.BadRequest(err.Error()))
			return
		}

		errs := validator.New().
			Required("name", req.Name).
			Required("email", req.Email).
			Email("email", req.Email).
			Required("password", req.Password).
			MinLen("password", req.Password, 8).
			Validate()

		if errs != nil {
			respond.Error(w, errutil.BadRequest(errs.Error()))
			return
		}

		respond.Created(w, map[string]string{"message": fmt.Sprintf("welcome, %s!", req.Name)})
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
