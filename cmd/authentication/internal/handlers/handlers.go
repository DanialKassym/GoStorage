package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	authentication "github.com/DanialKassym/GoStorage/cmd/authentication/internal/auth"
	"github.com/DanialKassym/GoStorage/cmd/authentication/internal/db"
	models "github.com/DanialKassym/GoStorage/cmd/authentication/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Authorize(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("here")
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 2)
	if err != nil {
		w.Write([]byte("couldnt hash the password"))
		return
	}
	if len(u.Username) > 3 && len(u.Password) > 3 {
		db.AddUser(u, string(hashedpass))
		tokenString, err := authentication.GenerateJWT(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println(tokenString)
		fmt.Println(authentication.ValidateJWT(tokenString))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Authorization", "Bearer "+tokenString)
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello to auth"))
	w.WriteHeader(200)
}
