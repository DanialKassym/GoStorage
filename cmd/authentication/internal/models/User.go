package models

type User struct {
	Username string `json:"Username" validate:"required,min=3"`
	Email string	`json:"Email" validate:"required,email"`
	Password string	`json:"Password" validate:"required,min=3"`
}
