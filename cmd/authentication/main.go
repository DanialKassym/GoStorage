package main

import (
	"fmt"

	authentication "github.com/DanialKassym/GoStorage/cmd/authentication/auth"
)

func main(){
	fmt.Println("hello")
	authentication.GenerateJWT()
	
}