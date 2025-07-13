package converter

import (
	"github.com/DanialKassym/GoStorage/api-gateway/internal/model"
	auth_proto "github.com/DanialKassym/protos/gen"
)

// ToModelFromProtoRegister converts a RegistrationResponse from gRPC to a RegistrationResponse model.
func ToModelFromProtoRegister(info *auth_proto.RegistrationResponse) *model.RegistrationResponse {
	return &model.RegistrationResponse{
		AccessToken: info.GetAccessToken(),
		UserID:      info.GetUserId(),
	}
}

// ToModelFromProtoLogin converts a LoginResponse from gRPC to a LoginResponse model.
func ToModelFromProtoLogin(info *auth_proto.LoginResponse) *model.LoginResponse {
	return &model.LoginResponse{
		AccessToken: info.GetAccessToken(),
		UserID:      info.GetUserId(),
	}
}