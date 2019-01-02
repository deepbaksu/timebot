package slack

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dl4b/timebot/timebot"
)

// CommandHandler handles slack slash command
//
// ENDPOINT /api/slack/command
//
// Example Usage
// /time 2018-12-31 21:40 PST
// => 2019-01-01 14:40 KST
func CommandHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	text := r.PostFormValue("text")

	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ret, err := timebot.ParseAndFlipTz(text)

	if err != nil {
		fmt.Fprintf(w, "%v is not a valid date time: %s", text, err)
		return
	}

	resp := Response{
		Text:         ret,
		ResponseType: InChannel,
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
