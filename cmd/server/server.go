package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dl4ab/timebot/api"
	"github.com/dl4ab/timebot/api/slack"
)

func mustLookupEnv(env string) string {
	ret, ok := os.LookupEnv(env)

	if !ok {
		log.Fatalf("Environment Variable %v is not available!\n", env)
	}

	return ret
}

func main() {
	// The env PORT is needed for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slackToken := mustLookupEnv("SLACK_SIGNING_SECRET")
	slackBotOAuthToken := mustLookupEnv("SLACK_BOT_OAUTH_TOKEN")

	app := slack.New(slackToken, slackBotOAuthToken)

	log.Printf("[MAIN] The server is running at 0.0.0.0:%v\n", port)
	log.Println("[MAIN]", http.ListenAndServe(":"+port, api.GetRouter(app)))
}
