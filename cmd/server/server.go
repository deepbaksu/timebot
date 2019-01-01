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
	log.Printf("The server is running at 0.0.0.0:%v\n", port)

	r := api.GetRouter()
	log.Println(http.ListenAndServe(":"+port, r))
}
