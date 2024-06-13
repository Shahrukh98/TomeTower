// pkg/util/util.go
package util

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// LoggingMiddleware logs details of each incoming HTTP request
func LoggingMiddleware(next *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request details
		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

		// Call the next handler (the underlying ServeMux)
		next.ServeHTTP(w, r)
	})
}

// JSONResponse is a utility function to send a JSON response.
func JSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// ParseJSONBody is a utility function to parse JSON request body.
func ParseJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dst)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}
	return nil
}

// HashPassword generates a bcrypt hash of the password with default cost
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if the provided password matches the hashed password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GetTimestamp returns the current time in a specific format.
func GetTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

// LogError logs an error with a custom message.
func LogError(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}