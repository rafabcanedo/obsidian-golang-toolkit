package health

import (
	"context"
	"encoding/json"
	"net/http"
)

type CheckFunc func(ctx context.Context) error

type Health struct {
	checks map[string]CheckFunc
}

func New() *Health {
	return &Health{
		checks: map[string]CheckFunc{},
	}
}

func (h *Health) AddCheck(name string, fn CheckFunc) *Health {
	h.checks[name] = fn
	return h
}

type response struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks,omitempty"`
}

func (h *Health) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		results := map[string]string{}
		degraded := false

		for name, check := range h.checks {
			if err := check(r.Context()); err != nil {
				results[name] = err.Error()
				degraded = true
			} else {
				results[name] = "ok"
			}
		}

		status := "ok"
		statusCode := http.StatusOK
		if degraded {
			status = "degraded"
			statusCode = http.StatusServiceUnavailable
		}

		resp := response{Status: status}
		if len(results) > 0 {
			resp.Checks = results
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(resp)
	})
}
