package slack

import (
	"log"
	"net/http"
)

// CommandHandler handles slack slash command
//
// ENDPOINT /api/slack/command
func CommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatalln(w.Write([]byte("NOT IMPLEMENTED")))
}
