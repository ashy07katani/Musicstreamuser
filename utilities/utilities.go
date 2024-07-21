package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytePassword), err
}
func ComparePassword(dbPassword string, loginPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(loginPassword))
	return err == nil
}

func FromJson[T any](body io.ReadCloser, user *T) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(user)
	return err
}

func GenerateJWTTokenUtil(jwtSecret string, email string) (string, error) {
	//creating claim
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "Musicstreamuser",
		"exp": time.Now().Add(time.Hour * 4).Unix(), // Expiration time
		"iat": time.Now().Unix(),
	})

	tokenString, err := claim.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func GenerateJWT(jwtSecret string, email string, rw http.ResponseWriter) {

	jwtToken, err := GenerateJWTTokenUtil(jwtSecret, email)
	response := make(map[string]interface{})
	if err != nil {
		response["status"] = "error"
		response["message"] = fmt.Sprintf("Error generating token %v", err)
		response["error"] = err.Error()
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
	} else {
		response["status"] = "success"
		response["message"] = "Token generated successfully"
		response["token"] = jwtToken
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(rw).Encode(response)
}
