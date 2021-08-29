package security

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/deepbaksu/timebot/config"
	"io/ioutil"
	"log"
	"net/http"
)

// checkMAC reports whether messageMAC is a valid HMAC tag for message.
func checkMAC(message, receivedMAC string, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	if _, err := mac.Write([]byte(message)); err != nil {
		log.Printf("mac.Write(%v) failed\n", message)
		return false
	}
	calculatedMAC := "v0=" + hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(calculatedMAC), []byte(receivedMAC))
}

// verify verifies the request is coming from Slack
//
// Read https://api.slack.com/docs/verifying-requests-from-slack
func verify(slackSigningToken []byte, timestamp, body, receivedMAC string) bool {
	message := "v0:" + timestamp + ":" + body
	return checkMAC(message, receivedMAC, slackSigningToken)
}

type Security struct {
	config *config.Config
}

// Verify is a wrapper around `verify`
func (s *Security) Verify(req *http.Request) bool {
	// do not consume req.body
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	bString := string(bodyBytes)
	ts := req.Header.Get("X-Slack-Request-Timestamp")
	eh := req.Header.Get("X-Slack-Signature")

	return verify([]byte(s.config.SlackToken), ts, bString, eh)
}

func ProvideSecurityService(cfg *config.Config) *Security {
	return &Security{cfg}
}
