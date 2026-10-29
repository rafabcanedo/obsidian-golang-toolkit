package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/auth"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

const testSecret = "test-secret-key"

func TestGenerate_Parse(t *testing.T) {
	claims := auth.Claims{
		Subject: "user123",
		Expiry:  time.Now().Add(time.Hour),
	}

	token, err := auth.Generate(testSecret, claims)
	testutil.NoError(t, err)

	got, err := auth.Parse(testSecret, token)
	testutil.NoError(t, err)
	testutil.Equal(t, got.Subject, claims.Subject)
}

func TestParse_invalidToken(t *testing.T) {
	_, err := auth.Parse(testSecret, "not.a.token")
	testutil.EqualError(t, err, auth.ErrInvalidToken.Error())
}

func TestParse_wrongSecret(t *testing.T) {
	claims := auth.Claims{Subject: "user123", Expiry: time.Now().Add(time.Hour)}
	token, _ := auth.Generate(testSecret, claims)

	_, err := auth.Parse("wrong-secret", token)
	testutil.EqualError(t, err, auth.ErrInvalidToken.Error())
}

func TestParse_expiredToken(t *testing.T) {
	claims := auth.Claims{Subject: "user123", Expiry: time.Now().Add(-time.Hour)}
	token, _ := auth.Generate(testSecret, claims)

	_, err := auth.Parse(testSecret, token)
	testutil.EqualError(t, err, auth.ErrExpiredToken.Error())
}

func TestMiddleware_valid(t *testing.T) {
	claims := auth.Claims{Subject: "user123", Expiry: time.Now().Add(time.Hour)}
	token, _ := auth.Generate(testSecret, claims)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	var got auth.Claims
	handler := auth.Middleware(testSecret)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got, _ = auth.GetClaims(r)
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rec, req)

	testutil.Equal(t, rec.Code, http.StatusOK)
	testutil.Equal(t, got.Subject, "user123")
}

func TestMiddleware_missingHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler := auth.Middleware(testSecret)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rec, req)
	testutil.Equal(t, rec.Code, http.StatusUnauthorized)
}

func TestMiddleware_invalidToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.here")
	rec := httptest.NewRecorder()

	handler := auth.Middleware(testSecret)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rec, req)
	testutil.Equal(t, rec.Code, http.StatusUnauthorized)
}
