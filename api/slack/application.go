package slack

import (
	"github.com/deepbaksu/timebot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type SlackRequestVerifier interface {
	Verify(r *http.Request) bool
}

// App manages Global Application State
type App struct {
	HttpClient           HttpClient
	MongoClient          *mongo.Client
	config               *config.Config
	slackRequestVerifier SlackRequestVerifier
}

func (app *App) GetOauthCollection() *mongo.Collection {
	return app.MongoClient.Database(app.config.MongoDBDATABASE).Collection("oauth")
}

func ProvideSlackApp(config *config.Config, httpClient HttpClient, mongoClient *mongo.Client, verifier SlackRequestVerifier) *App {
	return &App{
		HttpClient:           httpClient,
		MongoClient:          mongoClient,
		config:               config,
		slackRequestVerifier: verifier,
	}
}
