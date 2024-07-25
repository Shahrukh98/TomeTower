package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type UserClaims struct {
	ID   string
	Name string
	jwt.RegisteredClaims
}

func GetUserFromContext(r *http.Request) *UserClaims {
	if user, ok := r.Context().Value("props").(*UserClaims); ok {
		return user
	}
	return nil
}

func CreateToken(userID, name string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &UserClaims{
		ID:   userID,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*UserClaims, error) {
	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		user, err := VerifyToken(token)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Store user info in context
		ctx := context.WithValue(r.Context(), "props", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

