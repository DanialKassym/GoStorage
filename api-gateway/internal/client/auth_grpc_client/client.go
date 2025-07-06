package grpc_auth_client

import (
	"github.com/DanialKassym/GoStorage/api-gateway/internal/model"
)

type AuthClient interface {
	Login(request *model.LoginRequest) (*model.LoginResponse, error)

	Registration(request *model.RegistrationRequest) (*model.RegistrationResponse, error)

	ValidateToken(accessToken string) error
}
