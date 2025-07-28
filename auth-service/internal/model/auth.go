package model

type RegistrationRequest struct {
	Email    string 
	Password string 
	Username string 
}

type RegistrationResponse struct {
	AccessToken string 
	UserID      int64  
}

type LoginRequest struct {
	Email    string 
	Password string 
	Username string 
}

type LoginResponse struct {
	AccessToken string 
	UserID      int64  
}


