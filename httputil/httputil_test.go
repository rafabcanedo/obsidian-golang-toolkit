package httputil_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/errutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/httputil"
)

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	err := httputil.WriteJSON(w, http.StatusOK, map[string]string{"hello": "world"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Code != http.StatusOK {
		t.Errorf("status = %d; want %d", w.Code, http.StatusOK)
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q; want application/json",ct)
	}
}

func TestWriteError(t *testing.T) {
    w := httptest.NewRecorder()
    httputil.WriteError(w, errutil.NotFound("user not found"))

    if w.Code != http.StatusNotFound {
        t.Errorf("status = %d; want %d", w.Code, http.StatusNotFound)
    }
}

func TestWriteNoContent(t *testing.T) {
    w := httptest.NewRecorder()
    httputil.WriteNoContent(w)

    if w.Code != http.StatusNoContent {
        t.Errorf("status = %d; want %d", w.Code, http.StatusNoContent)
    }
}
