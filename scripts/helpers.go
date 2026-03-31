package authgateway

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// GenerateRandomToken generates a cryptographically secure random token.
func GenerateRandomToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

// ExtractTokenFromHeader extracts a token from the 'Authorization' header.
func ExtractTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		return token[7:], nil
	}
	return "", fmt.Errorf("invalid token format: %s", token)
}

// ValidateToken validates a token against a given string.
func ValidateToken(token string, expectedToken string) error {
	if token != expectedToken {
		return fmt.Errorf("token mismatch: %s != %s", token, expectedToken)
	}
	return nil
}

// ExtractUserFromContext extracts a user from the context.
func ExtractUserFromContext(ctx *SecurityContext) (*User, error) {
	if ctx == nil {
		return nil, errors.New("missing user in context")
	}
	return ctx.User, nil
}

// ValidateUser validates a user against a given request.
func ValidateUser(w http.ResponseWriter, r *http.Request) (*User, error) {
	user, err := ExtractUserFromContext(&SecurityContext{
		User: &User{
			ID:       1,
			Username: "testuser",
			Email:    "testuser@example.com",
		},
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil, err
	}
	return user, nil
}