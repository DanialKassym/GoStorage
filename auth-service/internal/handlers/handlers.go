package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	authentication "github.com/DanialKassym/GoStorage/auth-service/internal/auth"
	db "github.com/DanialKassym/GoStorage/auth-service/internal/db"
	models "github.com/DanialKassym/GoStorage/auth-service/internal/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var u models.LoginUser
	validate := validator.New(validator.WithRequiredStructEnabled())

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed on the '%s' tag\n", err.Field(), err.Tag())
		}
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	pass := db.Getuser(u.Username)
	error := bcrypt.CompareHashAndPassword([]byte(pass), []byte(u.Password))
	if error != nil {
		http.Error(w, "Credential Error", http.StatusBadRequest)
		return
	}

	tokenString, err := authentication.GenerateJWT(u.Password)
	if err != nil {
		http.Error(w, "Couldnt GenerateJWTtoken", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+tokenString)
}

func CheckJWT(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	fmt.Println("Raw header:", authHeader)

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
		return
	}

	tokenString := parts[1]
	fmt.Println("Token:", tokenString)

	valid := authentication.ValidateJWT(tokenString)
	if !valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return 
	} else {
		fmt.Fprintln(w, "Token is valid")
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	validate := validator.New(validator.WithRequiredStructEnabled())

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed on the '%s' tag\n", err.Field(), err.Tag())
		}
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 2)
	if err != nil {
		http.Error(w, "Couldnt hash the password", http.StatusBadRequest)
		return
	}
	db.AddUser(u, string(hashedpass))

	/*json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})*/

}

func Main(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello to auth"))
}
