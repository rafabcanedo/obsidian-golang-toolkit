package binding_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/binding"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestJSON(t *testing.T) {
	type payload struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	tests := []struct {
		name    string
		body    string
		wantErr bool
		want    payload
	}{
		{
			name: "valid JSON",
			body: `{"name":"Rafael","email":"rafael@email.com"}`,
			want: payload{Name: "Rafael", Email: "rafael@email.com"},
		},
		{
			name:    "invalid JSON",
			body:    `{invalid}`,
			wantErr: true,
		},
		{
			name:    "empty body",
			body:    ``,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			var got payload
			err := binding.JSON(req, &got)
			if tt.wantErr {
				testutil.Error(t, err)
				return
			}
			testutil.NoError(t, err)
			testutil.Equal(t, got, tt.want)
		})
	}
}

func TestQuery(t *testing.T) {
	type params struct {
		Page   int    `form:"page"`
		Limit  int    `form:"limit"`
		Search string `form:"search"`
		Active bool   `form:"active"`
	}

	tests := []struct {
		name    string
		query   string
		wantErr bool
		want    params
	}{
		{
			name:  "all fields",
			query: "page=2&limit=10&search=rafael&active=true",
			want:  params{Page: 2, Limit: 10, Search: "rafael", Active: true},
		},
		{
			name:  "missing fields use zero value",
			query: "page=1",
			want:  params{Page: 1},
		},
		{
			name:    "invalid int",
			query:   "page=abc",
			wantErr: true,
		},
		{
			name:    "invalid bool",
			query:   "active=maybe",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/?"+tt.query, nil)
			var got params
			err := binding.Query(req, &got)
			if tt.wantErr {
				testutil.Error(t, err)
				return
			}
			testutil.NoError(t, err)
			testutil.Equal(t, got, tt.want)
		})
	}
}
