package service

import (
	"Musicstreamuser/dto"
	models "Musicstreamuser/model"
	"Musicstreamuser/repository"
	"Musicstreamuser/utilities"
	"io"
	"log"
	"net/http"

	"github.com/jinzhu/copier"
)

type UserService struct {
	userRespository *repository.UserRepository
	jwtSecret       string
}

func NewService(userRepository *repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{userRespository: userRepository,
		jwtSecret: jwtSecret}
}

func (userService *UserService) LoginUser(r *http.Request, rw http.ResponseWriter) {
	loginRequest := &dto.LoginRequest{}
	err := utilities.FromJson(r.Body, loginRequest)
	if err != nil {
		log.Fatalf("Error converting input to valid user: %v", err)
	}
	//call the repository method to get the username with the password provided.
	loginResponse, err := userService.userRespository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		log.Fatalf("Error fetching the user with given email id: %v", err)
	}

	if utilities.ComparePassword(loginResponse.Password, loginRequest.Password) {
		log.Printf("user fetched: %v", loginResponse)
		utilities.GenerateJWT(userService.jwtSecret, loginRequest.Email, rw)
	} else {
		log.Printf("password Doesn't match: %v", err)
	}

}
func (userService *UserService) CreateUser(body io.ReadCloser) bool {

	//user := &models.User{}
	userdto := &dto.RegisterRequest{}
	err := utilities.FromJson(body, userdto)
	if err != nil {
		log.Fatalf("Error converting input to valid user: %v", err)
		return false
	}
	userdto.Password, err = utilities.HashPassword(userdto.Password)
	if err != nil {
		log.Fatalf("Error hashing the password: %v", err)
	}
	user := &models.User{}
	err = copier.Copy(user, userdto)
	if err != nil {
		log.Fatalf("Error converting DTO to model struct: %v", err)
		return false
	}
	err = userService.userRespository.CreateUser(user)
	if err != nil {
		log.Fatalf("Error creating a  user: %v", err)
		return false
	}
	//userRepository.CreateUser()
	return true
}
