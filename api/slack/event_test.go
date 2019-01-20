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

	app := New("")
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
