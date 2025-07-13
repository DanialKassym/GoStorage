package converter

import (
	"github.com/DanialKassym/GoStorage/api-gateway/internal/model"
	auth_proto "github.com/DanialKassym/protos/gen"
)

// FromModelToProtoRegister converts a RegistrationRequest model to a gRPC RegistrationRequest.
func FromModelToProtoRegister(info *model.RegistrationRequest) *auth_proto.RegistrationRequest {
	return &auth_proto.RegistrationRequest{
		Email:    info.Email,
		Password: info.Password,
		Name:     info.Username,
	}
}

// FromModelToProtoLogin converts a LoginRequest model to a gRPC LoginRequest.
func FromModelToProtoLogin(info *model.LoginRequest) *auth_proto.LoginRequest {
	return &auth_proto.LoginRequest{
		Email:    info.Email,
		Password: info.Password,
	}
}

// ToProtoValidateToken creates a gRPC ValidateTokenRequest from an access token string.
func ToProtoValidateToken(token string) *auth_proto.ValidateTokenRequest {
	return &auth_proto.ValidateTokenRequest{
		AccessToken: token,
	}
}
