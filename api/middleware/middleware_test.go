package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/middleware"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func okHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func TestLogger(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler := middleware.Logger(okHandler())
	handler.ServeHTTP(rr, req)

	testutil.Equal(t, rr.Code, http.StatusOK)
}

func TestRecovery(t *testing.T) {
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("something went wrong")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler := middleware.Recovery(panicHandler)
	handler.ServeHTTP(rr, req)

	testutil.Equal(t, rr.Code, http.StatusInternalServerError)
}

func TestRequestID(t *testing.T) {
	var capturedID string

	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedID = middleware.GetRequestID(r)
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler := middleware.RequestID(inner)
	handler.ServeHTTP(rr, req)

	testutil.Equal(t, rr.Code, http.StatusOK)
	testutil.Equal(t, len(capturedID) > 0, true)
	testutil.Equal(t, rr.Header().Get("X-Request-ID"), capturedID)
}

func TestCORS(t *testing.T) {
	opts := middleware.CORSOptions{
		AllowedOrigins: []string{"https://example.com"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	}

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Origin", "https://example.com")
	rr := httptest.NewRecorder()

	handler := middleware.CORS(opts)(okHandler())
	handler.ServeHTTP(rr, req)

	testutil.Equal(t, rr.Header().Get("Access-Control-Allow-Origin"), "https://example.com")
}

func TestCORS_preflight(t *testing.T) {
	opts := middleware.CORSOptions{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	}

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set("Origin", "https://anything.com")
	rr := httptest.NewRecorder()

	handler := middleware.CORS(opts)(okHandler())
	handler.ServeHTTP(rr, req)

	testutil.Equal(t, rr.Code, http.StatusNoContent)
}
