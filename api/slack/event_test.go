package slack

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TODO Clean up the test code
func TestEventHandler(t *testing.T) {
	body := `{
    "token": "Jhj5dZrVaK7ZwHHjRyZWjbDl",
    "challenge": "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P",
    "type": "url_verification"
}`
	req, _ := http.NewRequest("POST", "/api/slack/event", bytes.NewBufferString(body))

	app.TestMode = true

	handler := http.HandlerFunc(app.EventHandler)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Log(body)
		t.Fatalf("Status was not ok: %v", status)
	}

	expected := "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P"
	if rr.Body.String() != expected {
		t.Fatalf(`
Expected:
	%v
Received:
	%v`, expected, rr.Body.String())
	}

	badRequest, _ := http.NewRequest("POST", "/api/slack/event", bytes.NewBufferString(`{ "hello": "bad_request" }`))
	handler = http.HandlerFunc(app.EventHandler)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, badRequest)

	if status := rr.Code; status != http.StatusOK {
		t.Log(body)
		t.Fatalf("Status was not ok: %v", status)
	}

	expected = "Unknown Event: map[hello:bad_request]"
	if rr.Body.String() != expected {
		t.Fatalf(`
Expected:
	%v
Received:
	%v`, expected, rr.Body.String())
	}

}

func TestParseEvent(t *testing.T) {
	testCases := []struct {
		input    string
		expected interface{}
	}{
		{
			input: `{
    "token": "Jhj5dZrVaK7ZwHHjRyZWjbDl",
    "challenge": "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P",
    "type": "url_verification"
}`,
			expected: EventChallenge{
				Token:     "Jhj5dZrVaK7ZwHHjRyZWjbDl",
				Challenge: "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P",
				Type:      "url_verification",
			},
		},
		{
			// nolint: lll
			input: `{"token":"FWQ","team_id":"T8GMXUUFR","api_app_id":"AF5D1NN5D","event":{"client_msg_id":"15373339-8a9d-4d1c-884d-3acd0acb50a7","type":"message","text":"hello","user":"UAZBXKA74","ts":"1548086004.627000","thread_ts":"1548060038.614400","parent_user_id":"UBLV78JSU","channel":"C8FKWC50S","event_ts":"1548086004.627000","channel_type":"channel"},"type":"event_callback","event_id":"EvFK14FWQJ","event_time":1548086004,"authed_users":["UF95LP3NG"]}`,
			expected: EventMessage{
				Token:    "FWQ",
				TeamID:   "T8GMXUUFR",
				APIAppID: "AF5D1NN5D",
				Event: EventMessageType{
					ClientMsgID:  "15373339-8a9d-4d1c-884d-3acd0acb50a7",
					Type:         "message",
					Text:         "hello",
					User:         "UAZBXKA74",
					Ts:           "1548086004.627000",
					ThreadTs:     "1548060038.614400",
					ParentUserID: "UBLV78JSU",
					Channel:      "C8FKWC50S",
					EventTs:      "1548086004.627000",
					ChannelType:  "channel",
				},
				Type:        "event_callback",
				EventID:     "EvFK14FWQJ",
				EventTime:   1548086004,
				AuthedUsers: []string{"UF95LP3NG"},
			},
		},
	}

	for _, tCase := range testCases {
		v, err := ParseEvent([]byte(tCase.input))
		if err != nil {
			t.Fatal("Should not fail but failed")
		}

		if !reflect.DeepEqual(v, tCase.expected) {
			t.Fatalf(`
Expected:
%v
Received:
%v`, tCase.expected, v)
		}
	}
}
