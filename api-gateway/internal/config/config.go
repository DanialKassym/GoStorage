package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	envPath = ".env"
)

type Config struct {
	GRPCPort int
	TokenTTL time.Duration
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
	port,err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil{
		slog.Error("Invalid grpc port ")
	}
	token,err := strconv.Atoi(os.Getenv("TokenTTL"))
	if err != nil{
		slog.Error("Invalid tokenttl duration")
	}
	tokenttl := time.Duration(token) * time.Second
	return &Config{
		GRPCPort: port,
		TokenTTL: tokenttl,
	}
}
