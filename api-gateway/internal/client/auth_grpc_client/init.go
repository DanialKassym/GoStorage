package grpc_auth_client

import (
	"context"
	"time"

	"github.com/DanialKassym/GoStorage/api-gateway/internal/converter"
	"github.com/DanialKassym/GoStorage/api-gateway/internal/model"
	auth_proto "github.com/DanialKassym/protos/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	auth_client auth_proto.AuthClient
}

func NewGRPCClient(authAddr string) (*GRPCClient, error) {
	authConn, err := grpc.NewClient(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := &GRPCClient{
		auth_client: auth_proto.NewAuthClient(authConn),
	}

	return client, nil
}

func (g *GRPCClient) Login(request *model.LoginRequest) (*model.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := g.auth_client.Login(ctx, converter.FromModelToProtoLogin(request))
	if err != nil {

	}

	return converter.ToModelFromProtoLogin(result), nil
}

func (g *GRPCClient) Registration(
	request *model.RegistrationRequest) (*model.RegistrationResponse,string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := g.auth_client.Registration(ctx, converter.FromModelToProtoRegister(request))

	if err != nil {
	}

	return converter.ToModelFromProtoRegister(result), result.GetRefreshToken(), nil
}
func (g *GRPCClient) ValidateToken(accessToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	_, err := g.auth_client.ValidateToken(ctx,converter.ToProtoValidateToken(accessToken))
	if err != nil {

	}
	return nil
}
