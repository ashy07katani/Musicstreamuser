package repository

import (
	"Musicstreamuser/dto"
	models "Musicstreamuser/model"
	"database/sql"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) CreateUser(user *models.User) error {
	_, err := userRepository.db.Exec(InsertQuery, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.Birthday, user.Country)
	if err != nil {
		log.Fatalf("Error executing createuser query: %v", err)
	}
	return err
}

func (userRepository *UserRepository) GetUserByEmail(email string) (*dto.LoginResponse, error) {
	row := userRepository.db.QueryRow(GetUserByEmail, email)
	loginResponse := &dto.LoginResponse{}
	err := row.Scan(&loginResponse.Username, &loginResponse.Email, &loginResponse.Password, &loginResponse.FirstName, &loginResponse.LastName)
	return loginResponse, err
}
