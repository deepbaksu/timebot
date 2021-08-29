package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/deepbaksu/timebot/api/slack"
)

// ProvideRouter returns a root router for everything
func ProvideRouter(app *slack.App) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods("GET")
	// Handles a slash command "/time 2019-01-01 PST"
	r.HandleFunc("/api/slack/command", app.CommandHandler).Methods("POST")
	// Handles Slack Event Subscription
	r.HandleFunc("/api/slack/event", app.EventHandler).Methods("POST")

	// Handles Slack Oauth
	r.HandleFunc("/api/slack/oauth", app.OauthHandler).Methods("GET")
	return r
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
