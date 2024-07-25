package util

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ParseJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dst)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}
	return nil
}

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

func ClientErrorResponse(w http.ResponseWriter, status int, err error) {
	errorMessage := err.Error()

	LogError(err, errorMessage)
	response, err := json.Marshal(err)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func IsDateValid(date string) error {
	pattern := `^((\d{4}))[-](0[1-9]|1[012])[-]([0][1-9]|[12][0-9]|3[01])$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	validDate := re.MatchString(date)
	if validDate {
		return nil
	}
	return errors.New("invalid date")
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

func LogError(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
