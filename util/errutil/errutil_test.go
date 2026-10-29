package errutil_test

import (
	"net/http"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/util/errutil"
)

func TestNotFound(t *testing.T) {
	err := errutil.NotFound("user not found")

	if err.Error() != "user not found" {
		t.Errorf("Error() = %q; want %q", err.Error(), "user not found")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Errorf("StatusCode = %d; want %d", err.StatusCode, http.StatusNotFound)
	}
}

func TestBadRequest(t *testing.T) {
	err := errutil.BadRequest("invalid payload")

	if err.StatusCode != http.StatusBadRequest {
		t.Errorf("StatusCode = %d; want %d", err.StatusCode, http.StatusBadRequest)
	}
}

func TestUnauthorized(t *testing.T) {
	err := errutil.Unauthorized("invalid credentials")

	if err.StatusCode != http.StatusUnauthorized {
		t.Errorf("StatusCode = %d; want %d", err.StatusCode, http.StatusUnauthorized)
	}
}
