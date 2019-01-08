package slack

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CheckMAC reports whether messageMAC is a valid HMAC tag for message.
func checkMAC(message, receivedMAC string, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	if _, err := mac.Write([]byte(message)); err != nil {
		log.Printf("mac.Write(%v) failed\n", message)
		return false
	}
	calculatedMAC := "v0=" + hex.EncodeToString(mac.Sum(nil))

	return calculatedMAC == receivedMAC
}

// Verify verifies the request is coming from Slack
//
// Read https://api.slack.com/docs/verifying-requests-from-slack
func Verify(slackSigningToken []byte, timestamp, body, receivedMAC string) bool {
	message := "v0:" + timestamp + ":" + body
	return CheckMAC(message, receivedMAC, slackSigningToken)
}

// VerifyRequest is a wrapper around `Verify`
func VerifyRequest(req *http.Request, slackSigningToken []byte) bool {
	ts := req.Header.Get("X-Slack-Request-Timestamp")
	bCopy, _ := req.GetBody()
	b, _ := ioutil.ReadAll(bCopy)
	bString := string(b)
	eh := req.Header.Get("X-Slack-Signature")

	return Verify(slackSigningToken, ts, bString, eh)
}
