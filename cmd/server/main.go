package main

import (
	"log"
	"os"

	"github.com/saigenix/bidding-system/pkg/web"
)

func main() {
	router := web.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on :%s...", port)
	router.Run(":" + port)
}
