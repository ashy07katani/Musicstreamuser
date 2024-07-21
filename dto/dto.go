package dto

import "time"

type RegisterRequest struct {
	ID        int       `json:"userid"`
	Username  string    `json:"username"  validate:"required,min=3,max=50"`
	Email     string    `json:"email"  validate:"email"`
	Password  string    `json:"password"  validate:"required,min=6,max=20,must_contain_sepcial_character"`
	FirstName string    `json:"firstname" validate:"required,min=3,max=50"`
	LastName  string    `json:"lastname" validate:"required,min=3,max=50"`
	Birthday  time.Time `json:"birthday" validate:"required"`
	Country   string    `json:"country" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID        int    `json:"userid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"hashpassword"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
