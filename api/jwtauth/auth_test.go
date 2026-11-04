package jwtauth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/api/jwtauth"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

const testSecret = "test-secret-key"

func TestGenerate_Parse(t *testing.T) {
	claims := jwtauth.Claims{
		Subject: "user123",
		Expiry:  time.Now().Add(time.Hour),
	}

	token, err := jwtauth.Generate(testSecret, claims)
	testutil.NoError(t, err)

	got, err := jwtauth.Parse(testSecret, token)
	testutil.NoError(t, err)
	testutil.Equal(t, got.Subject, claims.Subject)
}

func TestParse_invalidToken(t *testing.T) {
	_, err := jwtauth.Parse(testSecret, "not.a.token")
	testutil.EqualError(t, err, jwtauth.ErrInvalidToken.Error())
}

func TestParse_wrongSecret(t *testing.T) {
	claims := jwtauth.Claims{Subject: "user123", Expiry: time.Now().Add(time.Hour)}
	token, _ := jwtauth.Generate(testSecret, claims)

	_, err := jwtauth.Parse("wrong-secret", token)
	testutil.EqualError(t, err, jwtauth.ErrInvalidToken.Error())
}

func TestParse_expiredToken(t *testing.T) {
	claims := jwtauth.Claims{Subject: "user123", Expiry: time.Now().Add(-time.Hour)}
	token, _ := jwtauth.Generate(testSecret, claims)

	_, err := jwtauth.Parse(testSecret, token)
	testutil.EqualError(t, err, jwtauth.ErrExpiredToken.Error())
}

func TestMiddleware_valid(t *testing.T) {
	claims := jwtauth.Claims{Subject: "user123", Expiry: time.Now().Add(time.Hour)}
	token, _ := jwtauth.Generate(testSecret, claims)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	var got jwtauth.Claims
	handler := jwtauth.Middleware(testSecret)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got, _ = jwtauth.GetClaims(r)
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rec, req)

	testutil.Equal(t, rec.Code, http.StatusOK)
	testutil.Equal(t, got.Subject, "user123")
}

func TestMiddleware_missingHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler := jwtauth.Middleware(testSecret)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rec, req)
	testutil.Equal(t, rec.Code, http.StatusUnauthorized)
}

func TestMiddleware_invalidToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.here")
	rec := httptest.NewRecorder()

	handler := jwtauth.Middleware(testSecret)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rec, req)
	testutil.Equal(t, rec.Code, http.StatusUnauthorized)
}
