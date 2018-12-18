package timebot

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	// KST = UTC + 9
	input := "2018-01-01 11:50 KST"
	expected := time.Date(2018, 1, 1, 11-9, 50, 0, 0, time.UTC)

	output, ok := ParseTime(input)

	if !ok || output != expected {
		t.Fatalf("Expected %v But received %v", expected, output)
	}

	// PST = UTC - 8
	input = "2018-12-18 14:17 PST"
	expected = time.Date(2018, 12, 18, 14+8, 17, 0, 0, time.UTC)

	output, ok = ParseTime(input)

	if !ok || output != expected {
		t.Fatalf("Expected %v But received %v", expected, output)
	}
}
