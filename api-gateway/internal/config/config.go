package config

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/joho/godotenv"
)

var (
	envPath = ".env"
)

type Config struct {
	HTTPAddr string
	GRPCAddr string
}

func LoadConfig() error {
	err := godotenv.Load(envPath)
	if err != nil {
		slog.Error("Error loading .env file")
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

func NewConfig() *Config {
	http_host := os.Getenv("HTTP_HOST")
	http_port := os.Getenv("HTTP_PORT")
	grpc_port := os.Getenv("GRPC_PORT")
	address := Address(http_host, http_port)
	return &Config{
		GRPCAddr: grpc_port,
		HTTPAddr: address,
	}
}

func Address(host string, port string) string {
	return net.JoinHostPort(host, port)
}
