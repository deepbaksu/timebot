package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dl4ab/timebot/timebot"
)

// EventHandler responds to the Slack Event
//
// When slack is first connected, it sends "Challenge"
// we need to return back the challenge code right away to be connected
func (a *App) EventHandler(w http.ResponseWriter, r *http.Request) {

	if !a.TestMode && !VerifyRequest(r, []byte(a.SigningToken)) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("EventHandler failed to read the body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	event, err := ParseEvent(body)

	if err != nil {
		log.Printf("ParseEvent(%s) returned an error", body)
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch v := event.(type) {
	case EventChallenge:
		// return a challenge
		w.Header().Set("Content-Type", "text/plain")
		log.Println(w.Write([]byte(v.Challenge)))
		return

	case EventMessage:
		// send ok right away
		w.WriteHeader(http.StatusOK)
		go checkMessageAndPostResponseIfInterested(a.BotOAuthAccessToken, v)
		return

	default:
		log.Print("===================================================")
		log.Printf(`Unknown event type is received:
%v`, v)
		log.Printf(`Original Body
%v`, string(body))
		log.Print("===================================================")
		fmt.Fprintf(w, "Unknown Event: %s", v)
		return
	}

}

// ParseEvent will try parsing slack events and return the first matching struct
//
// Possible event types
// - EventChallenge
//
// It will return map[string]interface{} if no type is matched
func ParseEvent(data []byte) (interface{}, error) {
	var anything map[string]interface{}
	err := json.Unmarshal(data, &anything)

	if err != nil {
		return nil, err
	}

	_, ok := anything["challenge"]

	if ok {
		return EventChallenge{
			Token:     anything["token"].(string),
			Challenge: anything["challenge"].(string),
			Type:      anything["type"].(string),
		}, nil
	}

	value, ok := anything["type"]

	if ok && value == "event_callback" {
		var event EventMessage
		err = json.Unmarshal(data, &event)
		return event, err
	}

	return anything, err
}

func checkMessageAndPostResponseIfInterested(token string, event EventMessage) {
	date, err := timebot.ExtractDateTime(event.Event.Text)

	if err != nil {
		// not interested in this message; so ignore
		return
	}

	flippedDate, err := timebot.ParseAndFlipTz(date)

	if err != nil {
		// something not right
		log.Printf(`timebot.ParseAndFlipTz returned an err:
%v`, err)
		return
	}

	message := ChatPostMessage{
		Token:   token,
		Channel: event.Event.Channel,
		Text:    fmt.Sprintf(`%v => %v`, date, flippedDate),
	}

	SendMessage(message)
}
