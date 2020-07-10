package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

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

// mongodb+srv://timebot-user:password@timebot.rgpco.gcp.mongodb.net/oauth_users
func obsfucate(mongoDbUri string) string {
	split := strings.Split(mongoDbUri, "://")

	if len(split) < 2 {
		return mongoDbUri
	}

	protocol := split[0]
	rest := split[1]

	split = strings.Split(rest, "@")

	if len(split) < 2 {
		return mongoDbUri
	}

	usernameAndPassword := split[0]
	rest = split[1]

	split = strings.Split(usernameAndPassword, ":")
	username := split[0]

	return fmt.Sprintf("%s://%s:xxxxxx@%s", protocol, username, rest)
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

	log.Printf("connecting to MONGODB_URI = %s", obsfucate(mongoDbUri))

	ctx, cancel1 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel1()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUri).SetRetryWrites(false).SetWriteConcern(writeconcern.New(writeconcern.WMajority())))
	fatalExitIfMongoError(err, mongoDbUri)

	ctx, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
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
