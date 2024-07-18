package models

import "time"

type User struct {
	ID        int       `json:"userid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Birthday  time.Time `json:"birthday"`
	Country   string    `json:"country"`
}
