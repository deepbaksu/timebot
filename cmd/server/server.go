package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dl4b/timebot/api"
)

func main() {
	// The env PORT is needed for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("[MAIN] The server is running at 0.0.0.0:%v\n", port)
	log.Println("[MAIN]", http.ListenAndServe(":"+port, api.GetRouter()))
}
