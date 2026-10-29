package ratelimit_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/ratelimit"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestAllow(t *testing.T) {
	limiter := ratelimit.New(3, time.Second)

	testutil.Equal(t, limiter.Allow("user1"), true)
	testutil.Equal(t, limiter.Allow("user1"), true)
	testutil.Equal(t, limiter.Allow("user1"), true)
	testutil.Equal(t, limiter.Allow("user1"), false)
}

func TestAllow_differentKeys(t *testing.T) {
	limiter := ratelimit.New(1, time.Second)

	testutil.Equal(t, limiter.Allow("user1"), true)
	testutil.Equal(t, limiter.Allow("user2"), true)
	testutil.Equal(t, limiter.Allow("user1"), false)
}

func TestMiddleware_allows(t *testing.T) {
	limiter := ratelimit.New(5, time.Second)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler := limiter.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rec, req)
	testutil.Equal(t, rec.Code, http.StatusOK)
}

func TestMiddleware_blocks(t *testing.T) {
	limiter := ratelimit.New(2, time.Second)

	handler := limiter.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = "1.2.3.4:1234"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		testutil.Equal(t, rec.Code, http.StatusOK)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "1.2.3.4:1234"
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	testutil.Equal(t, rec.Code, http.StatusTooManyRequests)
}

func TestMiddleware_xRealIP(t *testing.T) {
	limiter := ratelimit.New(1, time.Second)

	handler := limiter.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Real-IP", "10.0.0.1")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	testutil.Equal(t, rec.Code, http.StatusOK)

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Real-IP", "10.0.0.1")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	testutil.Equal(t, rec.Code, http.StatusTooManyRequests)
}
