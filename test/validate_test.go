package test

import (
	"Musicstreamuser/dto"
	"Musicstreamuser/validation"
	"testing"
	"time"
)

func TestValidateUserDto(test *testing.T) {
	// passInstance := dto.RegisterRequest{
	// 	ID:        1,
	// 	Username:  "john_doe",
	// 	Email:     "john.doe@example.com",
	// 	Password:  "securePassword123",
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// 	Birthday:  time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	// 	Country:   "USA",
	// }
	failInstance := dto.RegisterRequest{
		ID:        1,
		Username:  "john_doe",
		Email:     "john.doe@example.com",
		Password:  "Password1", // Invalid password: does not contain a special character
		FirstName: "John",
		LastName:  "Doe",
		Birthday:  time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Country:   "USA",
	}
	//err := validation.ValidateUser(&passInstance)
	err := validation.ValidateUser(&failInstance)
	if err != nil {
		test.Fatalf("Test Failed: %v", err)
	}

}
