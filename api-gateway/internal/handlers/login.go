package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	grpc_auth_client "github.com/DanialKassym/GoStorage/api-gateway/internal/client/auth_grpc_client"
	"github.com/DanialKassym/GoStorage/api-gateway/internal/model"
	"github.com/go-playground/validator/v10"
)

func Login(logger *slog.Logger, client *grpc_auth_client.GRPCClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Request model.LoginRequest
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		validate := validator.New(validator.WithRequiredStructEnabled())

		err := json.NewDecoder(r.Body).Decode(&Request)
		if err != nil {
			logger.Warn("Couldnt decode the Request")
			http.Error(w, "Couldnt decode the Request", http.StatusBadRequest)
			return
		}
		err = validate.Struct(Request)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				logger.Warn("Field '%s' failed on the '%s' tag\n", err.Field(), err.Tag())
			}
			http.Error(w, "Validation failed", http.StatusBadRequest)
			return
		}
		result, err := client.Login(ctx, &Request)
		if err != nil {
			logger.Error("failed to login")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Authorization", "Bearer " + result.AccessToken)
	}

}
