package slack

import "net/http"

// Verify verifies the request is coming from Slack
//
// Read https://api.slack.com/docs/verifying-requests-from-slack
func Verify(slackSigningToken, body []byte, timestamp int64, expectedHex []byte) bool {
	return false
}

// VerifyRequest is a wrapper around `Verify`
func VerifyRequest(req *http.Request, slackSigningToken []byte) bool {
	return false
}
