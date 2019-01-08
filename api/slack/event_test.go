package slack

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventHandler(t *testing.T) {
	body := `{
    "token": "Jhj5dZrVaK7ZwHHjRyZWjbDl",
    "challenge": "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P",
    "type": "url_verification"
}`
	req, _ := http.NewRequest("POST", "/api/slack/event", bytes.NewBufferString(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EventHandler)

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

}
