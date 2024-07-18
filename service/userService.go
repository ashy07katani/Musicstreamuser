package service

import (
	models "Musicstreamuser/model"
	"Musicstreamuser/repository"
	"Musicstreamuser/utilities"
	"io"
	"log"
)

type UserService struct {
	userRespository *repository.UserRepository
}

func NewService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRespository: userRepository}
}
func (userService *UserService) CreateUser(body io.ReadCloser) bool {

	user := &models.User{}
	err := utilities.FromJson(body, user)
	if err != nil {
		log.Fatalf("Error converting input to valid user: %v", err)
		return false
	}
	user.Password, err = utilities.HashPassword(user.Password)
	if err != nil {
		log.Fatalf("Error hashing the password: %v", err)
	}
	err = userService.userRespository.CreateUser(user)
	if err != nil {
		log.Fatalf("Error creating a  user: %v", err)
		return false
	}
	//userRepository.CreateUser()
	return true
}
