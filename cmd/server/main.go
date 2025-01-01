package main

import (
	"fmt"
	"os"
)

// ENV: PORT - port to run the server on

func main() {
	// router := web.SetupRouter()
	// log.Println("Starting server on :8080...")
	// router.Run(":8080")
	fmt.Println(os.Getenv("PORT")) // Works only with containerzied build
	fmt.Println("Hello, World!")
}
