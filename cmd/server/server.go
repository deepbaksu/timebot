package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dl4ab/timebot/api"
)

func main() {
	// The env PORT is needed for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slackToken, ok := os.LookupEnv("SLACK_SIGNING_SECRET")
	if !ok {
		log.Fatalf("Environment Variable SLACK_SIGNING_SECRET is not available!")
	}
	log.Printf("[MAIN] The server is running at 0.0.0.0:%v\n", port)
	log.Println("[MAIN]", http.ListenAndServe(":"+port, api.GetRouter(slackToken)))
}
