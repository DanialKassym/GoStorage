package main

import (
	"fmt"

	routes "github.com/DanialKassym/GoStorage/storage-service/internal/router"
)

func main() {
	fmt.Println("hello world")
	routes.InitRoutes()
}
