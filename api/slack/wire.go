//go:build wireinject
// +build wireinject

package slack

import (
	"github.com/deepbaksu/timebot/api/slack/security"
	"github.com/deepbaksu/timebot/config"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideTestApp(httpClient HttpClient, mongoClient *mongo.Client) *App {
	panic(wire.Build(
		ProvideSlackApp,
		config.TestSet,
		security.TestSet,
		wire.Bind(new(SlackRequestVerifier), new(*security.FakeSecurity)),
	))
}
