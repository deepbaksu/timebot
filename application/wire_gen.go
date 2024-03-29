// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"github.com/deepbaksu/timebot/api"
	"github.com/deepbaksu/timebot/api/slack"
	"github.com/deepbaksu/timebot/api/slack/security"
	"github.com/deepbaksu/timebot/config"
	"github.com/deepbaksu/timebot/database"
	"net/http"
)

// Injectors from wire.go:

func InitializeApplication() *Application {
	configConfig := config.ProvideProdConfig()
	httpClient := _wireClientValue
	client := database.ProvideMongoClient(configConfig)
	securitySecurity := security.ProvideSecurityService(configConfig)
	app := slack.ProvideSlackApp(configConfig, httpClient, client, securitySecurity)
	router := api.ProvideRouter(app)
	application := ProvideApplication(app, router)
	return application
}

var (
	_wireClientValue = http.DefaultClient
)
