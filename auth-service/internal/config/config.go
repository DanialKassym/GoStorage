package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var (
	envPath = ".env"
)

type Config struct {
	GRPCAddr string
	JWTKey string
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
	grpc_port := os.Getenv("AUTH_GRPC_PORT")
	jwt_key := os.Getenv("JWT_KEY")
	return &Config{
		GRPCAddr: grpc_port,
		JWTKey: jwt_key,
	}
}

