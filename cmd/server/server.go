package main

import (
	"github.com/deepbaksu/timebot/application"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	// The env PORT is needed for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := application.InitializeApplication().Start(port); err != nil {
		logrus.WithError(err).Panicf("application.Start(port: %s) has returned an error", port)
	}
}
