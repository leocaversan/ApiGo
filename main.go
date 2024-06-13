package main

import (
	"fmt"

	"go/src/routes"
)

func main() {
	fmt.Println("The server is running on port 8000!")
	routes.HandlerRequest()
}
