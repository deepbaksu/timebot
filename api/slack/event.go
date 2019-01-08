package slack

import (
	"encoding/json"
	"net/http"
)

// EventHandler responds to the Slack Event
func EventHandler(w http.ResponseWriter, r *http.Request) {
	var v EventChallenge
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, "Wrong body was received", http.StatusBadRequest)
		return
	}

	w.Write([]byte(v.Challenge))
}
