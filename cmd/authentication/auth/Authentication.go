package authentication

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWT() {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

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

	key = []byte(os.Getenv(("JWT_KEY")))
	t = jwt.New(jwt.SigningMethodHS256)
	s, err = t.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func ValidateJWT(tokenString string) bool {
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

	secret := []byte(os.Getenv(("JWT_KEY")))

	jwt, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	fmt.Println(jwt)

	if err != nil {
		fmt.Println("Invalid token:", err)
		return false
	} else {
		fmt.Println("Token is valid")
		return true

	}

}
