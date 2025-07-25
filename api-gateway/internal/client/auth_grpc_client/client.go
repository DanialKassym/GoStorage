package grpc_auth_client

import (
	"context"

	"github.com/DanialKassym/GoStorage/api-gateway/internal/model"
)

type AuthClient interface {
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse,error)

	Registration(ctx context.Context, request *model.RegistrationRequest) (*model.RegistrationResponse, string, error)

	ValidateToken(ctx context.Context, accessToken string) error
}
