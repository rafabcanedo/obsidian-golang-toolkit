package httpclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/httpclient"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testutil.Equal(t, r.Method, http.MethodGet)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := httpclient.New()
	resp, err := client.Get(context.Background(), server.URL)
	testutil.NoError(t, err)
	testutil.Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testutil.Equal(t, r.Method, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	client := httpclient.New()
	resp, err := client.Post(context.Background(), server.URL, []byte(`{"name":"Rafael"}`))
	testutil.NoError(t, err)
	testutil.Equal(t, resp.StatusCode, http.StatusCreated)
}

func TestWithHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testutil.Equal(t, r.Header.Get("X-Api-Key"), "secret")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := httpclient.New(httpclient.WithHeader("X-Api-Key", "secret"))
	resp, err := client.Get(context.Background(), server.URL)
	testutil.NoError(t, err)
	testutil.Equal(t, resp.StatusCode, http.StatusOK)
}

func TestWithTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := httpclient.New(httpclient.WithTimeout(10 * time.Millisecond))
	_, err := client.Get(context.Background(), server.URL)
	testutil.Error(t, err)
}

func TestWithRetry(t *testing.T) {
	count := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		if count < 3 {
			hj := w.(http.Hijacker)
			conn, _, _ := hj.Hijack()
			conn.Close()
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := httpclient.New(httpclient.WithRetry(3))
	resp, err := client.Get(context.Background(), server.URL)
	testutil.NoError(t, err)
	testutil.Equal(t, resp.StatusCode, http.StatusOK)
	testutil.Equal(t, count, 3)
}
