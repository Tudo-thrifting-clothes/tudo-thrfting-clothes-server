package main

import (
	"tudo-thrifting-server/routers"
)

func main() {
	// Set up the router
	r := routers.SetupRouter()

	// Run the Gin server on port 8080
	r.Run(":8080")
}
