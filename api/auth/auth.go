package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/contextutil"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type Claims struct {
	Subject string
	Expiry  time.Time
}

type jwtHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type jwtPayload struct {
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
	Iat int64  `json:"iat"`
}

const claimsKey contextutil.Key = "auth_claims"

func Generate(secret string, claims Claims) (string, error) {
	hJSON, err := json.Marshal(jwtHeader{Alg: "HS256", Typ: "JWT"})
	if err != nil {
		return "", err
	}

	pJSON, err := json.Marshal(jwtPayload{
		Sub: claims.Subject,
		Exp: claims.Expiry.Unix(),
		Iat: time.Now().Unix(),
	})
	if err != nil {
		return "", err
	}

	unsigned := base64url(hJSON) + "." + base64url(pJSON)
	return unsigned + "." + sign(secret, unsigned), nil
}

func Parse(secret, token string) (Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return Claims{}, ErrInvalidToken
	}

	unsigned := parts[0] + "." + parts[1]
	if !hmac.Equal([]byte(sign(secret, unsigned)), []byte(parts[2])) {
		return Claims{}, ErrInvalidToken
	}

	pJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return Claims{}, ErrInvalidToken
	}

	var p jwtPayload
	if err := json.Unmarshal(pJSON, &p); err != nil {
		return Claims{}, ErrInvalidToken
	}

	if time.Now().Unix() > p.Exp {
		return Claims{}, ErrExpiredToken
	}

	return Claims{
		Subject: p.Sub,
		Expiry:  time.Unix(p.Exp, 0),
	}, nil
}

func Middleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := Parse(secret, token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			r = r.WithContext(contextutil.WithValue(r.Context(), claimsKey, claims))
			next.ServeHTTP(w, r)
		})
	}
}

func GetClaims(r *http.Request) (Claims, bool) {
	return contextutil.Value[Claims](r.Context(), claimsKey)
}

func sign(secret, data string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func base64url(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}
