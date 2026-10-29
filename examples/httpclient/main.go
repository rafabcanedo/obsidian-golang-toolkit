package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/httpclient"
)

func main() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "method=%s auth=%s", r.Method, r.Header.Get("Authorization"))
	}))
	defer server.Close()

	client := httpclient.New(
		httpclient.WithTimeout(5*time.Second),
		httpclient.WithHeader("Authorization", "Bearer my-token"),
		httpclient.WithRetry(2),
	)

	resp, err := client.Get(context.Background(), server.URL+"/users")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("GET status:", resp.StatusCode)

	resp, err = client.Post(context.Background(), server.URL+"/users", []byte(`{"name":"Rafael"}`))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("POST status:", resp.StatusCode)
}
