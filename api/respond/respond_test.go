package respond_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/respond"
	"github.com/rafabcanedo/obsidian-golang-toolkit/pagination"
	"github.com/rafabcanedo/obsidian-golang-toolkit/util/errutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestOK(t *testing.T) {
	rec := httptest.NewRecorder()
	respond.OK(rec, map[string]string{"name": "Rafael"})

	testutil.Equal(t, rec.Code, http.StatusOK)
	testutil.Equal(t, rec.Header().Get("Content-Type"), "application/json")

	var body map[string]any
	json.NewDecoder(rec.Body).Decode(&body)
	_, hasData := body["data"]
	testutil.Equal(t, hasData, true)
}

func TestCreated(t *testing.T) {
	rec := httptest.NewRecorder()
	respond.Created(rec, map[string]string{"id": "1"})

	testutil.Equal(t, rec.Code, http.StatusCreated)
}

func TestNoContent(t *testing.T) {
	rec := httptest.NewRecorder()
	respond.NoContent(rec)

	testutil.Equal(t, rec.Code, http.StatusNoContent)
	testutil.Equal(t, rec.Body.Len(), 0)
}

func TestError_appError(t *testing.T) {
	rec := httptest.NewRecorder()
	respond.Error(rec, errutil.NotFound("user not found"))

	testutil.Equal(t, rec.Code, http.StatusNotFound)

	var body map[string]any
	json.NewDecoder(rec.Body).Decode(&body)
	testutil.Equal(t, body["error"], "user not found")
}

func TestError_genericError(t *testing.T) {
	rec := httptest.NewRecorder()
	respond.Error(rec, errors.New("something went wrong"))

	testutil.Equal(t, rec.Code, http.StatusInternalServerError)
}

func TestPaginated(t *testing.T) {
	rec := httptest.NewRecorder()
	meta := pagination.New(1, 10).NewMeta(100)
	respond.Paginated(rec, []string{"a", "b"}, meta)

	testutil.Equal(t, rec.Code, http.StatusOK)

	var body map[string]any
	json.NewDecoder(rec.Body).Decode(&body)
	_, hasData := body["data"]
	_, hasMeta := body["meta"]
	testutil.Equal(t, hasData, true)
	testutil.Equal(t, hasMeta, true)
}
