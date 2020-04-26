package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"time"

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
	slackClientId := mustLookupEnv("SLACK_CLIENT_ID")
	slackClientSecret := mustLookupEnv("SLACK_CLIENT_SECRET")

	mongoDbUri := mustLookupEnv("MONGODB_URI")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUri))
	fatalExitIfMongoError(err, mongoDbUri)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Ping(ctx, readpref.Primary())
	fatalExitIfMongoError(err, mongoDbUri)

	httpClientImpl := slack.HttpClientImpl{Client: http.DefaultClient}
	app := slack.New(slackToken, slackBotOAuthToken, slackClientId, slackClientSecret, &httpClientImpl, mongoClient)

	log.Printf("[MAIN] The server is running at 0.0.0.0:%v", port)
	log.Println("[MAIN]", http.ListenAndServe(":"+port, api.GetRouter(app)))
}

func fatalExitIfMongoError(err error, mongoDbUri string) {
	if err != nil {
		log.Fatalf("Failed to connect MongoDB. Please check $MONGODB_URI(%v) is a correct MongoDB URI (err => %v)", mongoDbUri, err)
	}
}
