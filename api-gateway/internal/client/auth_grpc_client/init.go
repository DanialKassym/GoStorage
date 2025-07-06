package grpc_auth_client

import (
	"log"
	proto_auth "github.com/DanialKassym/GoStorage/tree/dev/auth-service/pkg/gen"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	Client  
}

func NewGRPCClient(inventoryAddr, orderAddr, statsAddr string) (*GRPCClient, error) {
	inventoryConn, err := grpc.Dial(inventoryAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to inventory service: %v", err)
		return nil, err
	}

	orderConn, err := grpc.Dial(orderAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
		return nil, err
	}
	statsConn, err := grpc.Dial(statsAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to statistics service: %v", err)
		return nil, err
	}

	client := &GRPCClient{
		InventoryClient:  proto.NewInventoryServiceClient(inventoryConn),
		OrderClient:      proto.NewOrderServiceClient(orderConn),
		StatisticsClient: proto.NewStatisticsServiceClient(statsConn),
	}

	return client, nil
}
