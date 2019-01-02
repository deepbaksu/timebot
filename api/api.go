package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dl4ab/timebot/api/slack"
)

// GetRouter returns a root router for everything
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods("GET")
	r.HandleFunc("/api/slack/command", slack.CommandHandler).Methods("POST")
	return r
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
