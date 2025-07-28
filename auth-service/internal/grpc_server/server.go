package grpc_server

import (
	"context"

	"github.com/DanialKassym/GoStorage/auth-service/internal/model"
)

type AuthServer interface {
	Registration(ctx context.Context,
		registrationRequest *model.RegistrationRequest) (*model.RegistrationResponse, error)

	Login(ctx context.Context, loginRequest *model.LoginRequest) (*model.LoginResponse, error)
	
	ValidateToken(ctx context.Context, accessToken string) (bool, error)
}
