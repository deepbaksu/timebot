package timebot

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	/////////////////////////////////
	// TEST CASES
	/////////////////////////////////
	var testCasesParseTime = []struct {
		input    string
		expected time.Time
		ok       bool
	}{
		{
			// KST = UTC + 9
			input:    "2018-01-01 11:50 KST",
			expected: time.Date(2018, 1, 1, 11-9, 50, 0, 0, time.UTC),
			ok:       true,
		},
		{
			// PST = UTC - 8
			input:    "2018-12-18 14:17 PST",
			expected: time.Date(2018, 12, 18, 14+8, 17, 0, 0, time.UTC),
			ok:       true,
		},
		{
			// PDT = UTC -7
			input:    "2018-08-13 20:00 PDT",
			expected: time.Date(2018, 8, 13, 20+7, 0, 0, 0, time.UTC),
			ok:       true,
		},
		{
			// Test Invalid PST/PDT case
			// 2018-12-21 19:34 PDT is not a valid date
			input: "2018-12-21 19:34 PDT",
			ok:    false,
		},
		{
			// "2018-08-13 19:34 PST" is not a PST
			input: "2018-08-13 19:34 PST",
			ok:    false,
		},
	}

	//////////////////////////////////
	// REMOVE THIS LINE
	//////////////////////////////////
	// t.Skip("[TEST SKIP] PLEASE REMOVE THIS")
	for _, testCase := range testCasesParseTime {

		output, ok := ParseTime(testCase.input)

		switch {

		case testCase.ok:
			// test should succeed
			if !ok || !output.Equal(testCase.expected) {
				t.Fatalf("Expected %v But received %v", testCase.expected, output)
			}

		default:
			// test should fail
			if ok {
				t.Fatal("Test should fail but succeeded")
			}
		}

	}
}

func TestParseAndFlipTz(t *testing.T) {
	/////////////////////////////////
	// TEST CASES
	/////////////////////////////////
	var testCases = []struct {
		input    string
		expected string

		// if false, it should fail
		ok bool
	}{
		{
			// Test PDT -> KST
			input:    "2018-08-13 20:00 PDT",
			expected: "2018-08-14 12:00 KST",
			ok:       true,
		},
		{
			// Test KST -> PDT
			input:    "2018-08-14 12:00 KST",
			expected: "2018-08-13 20:00 PDT",
			ok:       true,
		},
		{
			// Test KST -> PST
			input:    "2018-12-21 19:34 PST",
			expected: "2018-12-22 12:34 KST",
			ok:       true,
		},
		{
			// Test Invalid PST/PDT case
			// 2018-12-21 19:34 PDT is not a valid date
			input: "2018-12-21 19:34 PDT",
			ok:    false,
		},
		{

			// "2018-08-13 19:34 PST" is not a PST
			input: "2018-08-13 19:34 PST",
			ok:    false,
		},
	}
	//////////////////////////////////
	// REMOVE THIS LINE
	//////////////////////////////////
	// t.Skip("[TEST SKIP] PLEASE REMOVE THIS")

	for _, testCase := range testCases {
		ret, err := ParseAndFlipTz(testCase.input)

		switch {
		case testCase.ok:
			// test should succeed
			if err != nil {
				t.Fatal("Test should succeed but failed")
			}

			if ret != testCase.expected {
				t.Fatalf("Expected %v but received %v", testCase.expected, ret)
			}
		default:
			// test should fail
			if err == nil {
				t.Fatal("Test should fail but succeeded")
			}

		}
	}
}
