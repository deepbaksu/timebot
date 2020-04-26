package slack

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HttpClientImpl struct {
	Client *http.Client
}

func (c *HttpClientImpl) Do(r *http.Request) (*http.Response, error) {
	return c.Client.Do(r)
}

// App manages Global Application State
type App struct {
	SigningToken        string
	BotOAuthAccessToken string
	SlackClientId       string
	SlackClientSecret   string
	HttpClient          HttpClient
	// if it's true, the request verification does not happen
	TestMode    bool
	MongoClient *mongo.Client
}

func (app *App) GetOauthCollection() *mongo.Collection {
	return app.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("oauth")
}

// New creates a slack API application
func New(slackSigningToken, botOAuthAccessToken, slackClientId, slackClientSecret string, httpClient HttpClient, mongoClient *mongo.Client) App {

	return App{
		SigningToken:        slackSigningToken,
		BotOAuthAccessToken: botOAuthAccessToken,

		SlackClientId:     slackClientId,
		SlackClientSecret: slackClientSecret,
		HttpClient:        httpClient,
		MongoClient:       mongoClient,

		TestMode: false,
	}
}
