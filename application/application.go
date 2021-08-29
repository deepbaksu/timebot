package application

import (
	"github.com/deepbaksu/timebot/api/slack"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Application struct {
	slackApp *slack.App
	router   *mux.Router
}

func (a *Application) Start(port string) error {
	logrus.Infof("application is starting at 0.0.0.0:%s", port)
	return http.ListenAndServe(":"+port, a.router)
}

func ProvideApplication(slackApp *slack.App, router *mux.Router) *Application {
	return &Application{
		slackApp: slackApp,
		router:   router,
	}
}
