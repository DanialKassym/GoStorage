package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	users "github.com/DanialKassym/GoStorage/cmd/rest-server/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func GetAllUsers() []users.User {
	// uncomment if you wish to run locally

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

	rows, err := dbpool.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed select: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var ret []users.User
	for rows.Next() {
		var u users.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
			fmt.Fprintf(os.Stderr, "Couldnt retrieve the object: %v\n", err)
			os.Exit(1)
		}
		ret = append(ret, u)
	}
	fmt.Println("here")
	fmt.Println(ret)

	return ret
}
