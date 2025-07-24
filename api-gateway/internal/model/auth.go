package model

type RegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	Username string `json:"name" validate:"required,min=5,max=15"`
}

type RegistrationResponse struct {
	AccessToken string `json:"accessToken"`
	UserID      int64  `json:"userId"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	Username string `json:"Username" validate:"required,min=5,max=15"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	UserID      int64  `json:"userId"`
}


