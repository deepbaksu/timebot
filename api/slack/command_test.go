package slack

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// nolint gochecknoglobals
var app App = New("abc", "", "", "")

func TestBasicRequest(t *testing.T) {
	body := url.Values{"text": {"2018-12-31 22:19 PST"}}.Encode()
	req, err := http.NewRequest("POST", "/api/slack/command", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal("Failed to build a request")
	}

	app.TestMode = true

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CommandHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Log(body)
		t.Fatalf("Status was not ok: %v", status)
	}

	expected := `{"text":"2019-01-01 15:19 KST","response_type":"ephemeral"}`

	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Fatalf("\nExpected:\n%s\nReceived:\n%s", expected, rr.Body.String())
	}
}

func must(req *http.Request, err error) *http.Request {
	return req
}

func TestBadRequest(t *testing.T) {
	testCases := []struct {
		request *http.Request
	}{
		{
			// no body
			request: must(http.NewRequest("POST", "/api/slack/command", nil)),
		},
		{
			// no "text" property
			request: must(http.NewRequest("POST", "/api/slack/command", bytes.NewBufferString(""))),
		},
	}

	req, _ := http.NewRequest("POST", "/api/slack/command", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	app.TestMode = true

	for _, testCase := range testCases {
		req = testCase.request
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.CommandHandler)

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("Expected a bad request but received %v", rr.Code)
		}
	}

}

func TestInvalidDateFormat(t *testing.T) {
	body := `text="not a date"`
	req, _ := http.NewRequest("POST", "/api/slack/command", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	app.TestMode = true

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CommandHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected a StatusOK but received %v", rr.Code)
	}

	expected := `"not a date" is not a valid date time: "not a date" does not contain PST/PDT/KST`
	if rr.Body.String() != expected {
		t.Fatalf(`
Expected
  %v
but received
  %v`, expected, rr.Body)
	}
}
