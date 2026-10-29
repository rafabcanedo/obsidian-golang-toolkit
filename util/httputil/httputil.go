package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/rafabcanedo/obsidian-golang-toolkit/errutil"
)

type errorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, err *errutil.AppError) error {
	return WriteJSON(w, err.StatusCode, errorResponse{Error: err.Message})
}

func WriteNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
