package slack

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestBasicRequest(t *testing.T) {
	body := url.Values{"text": {"2018-12-31 22:19 PST"}}.Encode()
	req, err := http.NewRequest("POST", "/api/slack/command", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal("Failed to build a request")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CommandHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Log(body)
		t.Fatalf("Status was not ok: %v", status)
	}

	expected := `{"text":"2019-01-01 15:19 KST","response_type":"in_channel"}`

	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Fatalf("\nExpected:\n%s\nReceived:\n%s", expected, rr.Body.String())
	}
}
