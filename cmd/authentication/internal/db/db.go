package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DanialKassym/GoStorage/cmd/authentication/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func AddUser(user models.User, hashedpass string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current directory: ", err)
		os.Exit(1)
	}

	envFilePath := filepath.Join(cwd, ".env")

	err = godotenv.Load(envFilePath)
	if err != nil {
		fmt.Println("error loading .env: ", err)
		os.Exit(1)
	}

	db := os.Getenv("DB_URL")

	dbpool, err := pgxpool.New(context.Background(), db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "INSERT INTO users (name, email, password) VALUES ($1, $2 , $3);",user.Username, user.Email, hashedpass)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed : %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

}
