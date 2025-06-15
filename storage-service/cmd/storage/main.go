package main

import (
	"fmt"

	routes "github.com/DanialKassym/GoStorage/cmd/rest-server/internal/router"
)

func main() {
	fmt.Println("hello world")
	routes.InitRoutes()
}
