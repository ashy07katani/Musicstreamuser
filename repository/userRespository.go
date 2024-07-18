package repository

import (
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
