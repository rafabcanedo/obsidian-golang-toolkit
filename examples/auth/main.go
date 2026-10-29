package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/auth"
)

func main() {
	secret := "my-secret-key"

	token, err := auth.Generate(secret, auth.Claims{
		Subject: "user-123",
		Expiry:  time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("token:", token)

	claims, err := auth.Parse(secret, token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("subject:", claims.Subject)

	protected := auth.Middleware(secret)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := auth.GetClaims(r)
		fmt.Fprintf(w, "hello, %s!", c.Subject)
	}))

	http.Handle("/protected", protected)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
