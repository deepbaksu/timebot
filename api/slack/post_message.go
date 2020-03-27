package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// ChatPostMessage is a struct sent to Slack
type ChatPostMessage struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Text    string `json:"text"`

	Attachments []interface{} `json:"attchments,omitempty"`
	ThreadTs    string        `json:"thread_ts,omitempty"`
	Mrkdown     bool          `json:"mrkdown,omitempty"`
	AsUser      bool          `json:"as_user,omitempty"`
}

// SendMessage sends to slack
//
// Documentation: https://api.slack.com/methods/chat.postMessage
func SendMessage(postMessage ChatPostMessage) {
	body, err := json.Marshal(postMessage)
	if err != nil {
		log.Printf(`Err while sending message:
%v`, postMessage)
		return
	}

	// nolint: lll
	req, err := http.NewRequest(
		http.MethodPost,
		"https://slack.com/api/chat.postMessage",
		bytes.NewBuffer(body),
	)

	// "application/json; charset=utf-8"
	if err != nil {
		log.Println(`Unable to create a new request while sending a message`)
		return
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", "Bearer "+postMessage.Token)

	log.Printf("Sending a message: %v", req)

	resp, _ := http.DefaultClient.Do(req)

	var slackResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&slackResponse)

	if err != nil {
		log.Println("Unable to decode the slack response")
		return
	}

	log.Printf(`Sent a message. Here's the response from Slack:
%v`, slackResponse)

}
