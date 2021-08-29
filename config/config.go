package config

import (
	"log"
	"os"
)

type Config struct {
	SlackToken         string
	SlackBotOAuthToken string
	SlackClientID      string
	SlackClientSecret  string
	MongoDBURI         string
	MongoDBDATABASE    string
}

func mustLookupEnv(env string) string {
	ret, ok := os.LookupEnv(env)

	if !ok {
		log.Fatalf("Environment Variable %v is not available!\n", env)
	}

	return ret
}

func ProvideProdConfig() *Config {
	return &Config{
		SlackToken:         mustLookupEnv("SLACK_SIGNING_SECRET"),
		SlackBotOAuthToken: mustLookupEnv("SLACK_BOT_OAUTH_TOKEN"),
		SlackClientID:      mustLookupEnv("SLACK_CLIENT_ID"),
		SlackClientSecret:  mustLookupEnv("SLACK_CLIENT_SECRET"),
		MongoDBURI:         mustLookupEnv("MONGODB_URI"),
		MongoDBDATABASE:    mustLookupEnv("MONGODB_DATABASE"),
	}
}

func ProvideTestConfig() *Config {
	return &Config{
		SlackToken:         "SLACK_SIGNING_SECRET",
		SlackBotOAuthToken: "SLACK_BOT_OAUTH_TOKEN",
		SlackClientID:      "SLACK_CLIENT_ID",
		SlackClientSecret:  "SLACK_CLIENT_SECRET",
		MongoDBURI:         "MONGODB_URI",
		MongoDBDATABASE:    "MONGODB_DATABASE",
	}
}
