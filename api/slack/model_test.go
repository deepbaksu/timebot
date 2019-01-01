package slack

import (
	"encoding/json"
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
