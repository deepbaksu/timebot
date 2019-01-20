package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dl4ab/timebot/api/slack"
)

// GetRouter returns a root router for everything
func GetRouter(slackSigningToken string) *mux.Router {
	app := slack.New(slackSigningToken)

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods("GET")
	// Handles a slash command "/time 2019-01-01 PST"
	r.HandleFunc("/api/slack/command", app.CommandHandler).Methods("POST")
	// Handles Slack Event Subscription
	r.HandleFunc("/api/slack/event", app.EventHandler).Methods("POST")
	return r
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
