package slack

import (
	"net/http"
)

// CommandHandler handles slack slash command
//
// ENDPOINT /api/slack/command
func CommandHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NOT IMPLEMENTED"))
}
