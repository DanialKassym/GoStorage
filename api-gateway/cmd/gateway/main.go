package main

import (
	"fmt"

	routes "github.com/DanialKassym/GoStorage/api-gateway/internal/router"
)

func main() {
	fmt.Println("hello world")
	routes.InitRoutes()
}
