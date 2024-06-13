package main

import (
	"fmt"

	"go/routes"
)

func main() {
	fmt.Println("The server is running on port 8000!")
	routes.HandlerRequest()
}
