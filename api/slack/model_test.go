package slack

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestResponseModelMarshall(t *testing.T) {
	testCases := []struct {
		input    Response
		expected string
	}{
		{
			input: Response{
				Text:         "hello",
				ResponseType: InChannel,
			},
			expected: `{"text":"hello","response_type":"in_channel"}`,
		},

		{
			input: Response{
				Text: "hello",
			},
			expected: `{"text":"hello"}`,
		},
	}

	for _, testCase := range testCases {

		r, err := json.Marshal(testCase.input)

		if err != nil {
			t.Fatal("Marshalling failed")
		}

		if string(r) != testCase.expected {
			t.Fatalf("Expected %s but received %s", testCase.expected, r)
		}
	}

}

func TestEventChallengeMarshallAndUnmarshal(t *testing.T) {
	// nolint gosec
	Token := "Jhj5dZrVaK7ZwHHjRyZWjbDl"
	Challenge := "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P"
	Type := "url_verification"

	eventChallenge := EventChallenge{
		Token,
		Challenge,
		Type,
	}

	j, err := json.Marshal(eventChallenge)

	if err != nil {
		t.Fatalf("json.Marshal(%v) returned err: %v", eventChallenge, err)
	}

	expected := fmt.Sprintf(`{"token":"%s","challenge":"%s","type":"%s"}`, Token, Challenge, Type)

	if string(j) != expected {
		t.Fatalf(`
Expected:
	%s
Received:
	%s\n`, expected, j)
	}

	var e EventChallenge
	err = json.Unmarshal([]byte(expected), &e)
	if err != nil {
		t.Fatalf("Failed to unmarshal:\n%s", expected)
	}

	if !reflect.DeepEqual(e, eventChallenge) {
		t.Fatalf(`
Expected:
	%v
Received:
	%v
`, eventChallenge, e)
	}

}
