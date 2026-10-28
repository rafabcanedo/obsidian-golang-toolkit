package health_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/health"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestHandler_noChecks(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	health.New().Handler().ServeHTTP(rec, req)

	testutil.Equal(t, rec.Code, http.StatusOK)
}

func TestHandler_allChecksPass(t *testing.T) {
	h := health.New().
		AddCheck("db", func(ctx context.Context) error { return nil })

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	h.Handler().ServeHTTP(rec, req)

	testutil.Equal(t, rec.Code, http.StatusOK)

	var resp map[string]any
	json.NewDecoder(rec.Body).Decode(&resp)
	testutil.Equal(t, resp["status"], "ok")
}

func TestHandler_checkFails(t *testing.T) {
	h := health.New().
		AddCheck("db", func(ctx context.Context) error {
			return errors.New("connection refused")
		})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	h.Handler().ServeHTTP(rec, req)

	testutil.Equal(t, rec.Code, http.StatusServiceUnavailable)

	var resp map[string]any
	json.NewDecoder(rec.Body).Decode(&resp)
	testutil.Equal(t, resp["status"], "degraded")
}
