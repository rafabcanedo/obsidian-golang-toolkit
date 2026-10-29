package respond

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rafabcanedo/obsidian-golang-toolkit/errutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/pagination"
)

type envelope struct {
	Data  any              `json:"data,omitempty"`
	Error string           `json:"error,omitempty"`
	Meta  *pagination.Meta `json:"meta,omitempty"`
}

func write(w http.ResponseWriter, status int, env envelope) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(env)
}

func OK(w http.ResponseWriter, data any) {
	write(w, http.StatusOK, envelope{Data: data})
}

func Created(w http.ResponseWriter, data any) {
	write(w, http.StatusCreated, envelope{Data: data})
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func Error(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	var appErr *errutil.AppError
	if errors.As(err, &appErr) {
		status = appErr.StatusCode
	}
	write(w, status, envelope{Error: err.Error()})
}

func Paginated(w http.ResponseWriter, data any, meta pagination.Meta) {
	write(w, http.StatusOK, envelope{Data: data, Meta: &meta})
}
