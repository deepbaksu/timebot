//go:build wireinject
// +build wireinject

package application

import (
	"github.com/deepbaksu/timebot/api"
	"github.com/deepbaksu/timebot/api/slack"
	"github.com/deepbaksu/timebot/config"
	"github.com/deepbaksu/timebot/database"
	"github.com/google/wire"
)

func InitializeApplication() *Application {
	panic(wire.Build(ProdSet, config.ProdSet, database.ProdSet, api.ProdSet, slack.ProdSet))
}
