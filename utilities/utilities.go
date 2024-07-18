package utilities

import (
	models "Musicstreamuser/model"
	"encoding/json"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	return string(bytePassword), err
}

func FromJson(body io.ReadCloser, user *models.User) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(user)
	return err
}
