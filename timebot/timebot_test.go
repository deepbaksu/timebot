package timebot

import (
	"testing"
	"time"
)

func TestCheckDaylightSavingZone(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"2018-12-30 11:30 KST", true},
		{"2018-12-30 11:11 PDT", false},
	}

	for _, testCase := range testCases {
		output := CheckDaylightSavingZone(testCase.input)
		if output != testCase.expected {
			t.Fatalf(`
Expected
  %v
But received
  %v`, testCase.expected, output)
		}
	}
}

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
		{
			// invalid date should return false
			input: "invalid date",
			ok:    false,
		},
	}

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
		{
			input: "no date time",
			ok:    false,
		},
	}
	//////////////////////////////////
	// REMOVE THIS LINE
	//////////////////////////////////

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

func TestExtractDateTime(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      bool
	}{
		{
			input:    "이번 미팅은 2019-01-04 21:51 KST 에 하겠습니다",
			expected: "2019-01-04 21:51 KST",
		},
		{
			// empty string should fail
			input: "",
			err:   true,
		},
		{
			// if it doesn't contain datetime, err
			input: "이번 미팅은 에 하겠습니다",
			err:   true,
		},
		{
			input:    "ㄴㅇ이ㄱ자ㅂㄴㅇㅣㅇㅂ자ㅇㄴㅇㅣㅂ자 2019-01-04 21:52 PST",
			expected: "2019-01-04 21:52 PST",
		},
	}

	for _, testCase := range testCases {
		output, err := ExtractDateTime(testCase.input)

		switch testCase.err {

		case true:

			if err == nil {

				// "should fail" case but succeeded
				t.Fatalf(`
test should fail but did not return error

Input:
	%v
`, testCase.input)
			}

		case false:

			if err != nil {
				// "success" case but failed
				t.Fatalf(`
test should not fail but failed

Input:
	%v
Error:
	%v
`, testCase.input, err)

			}

			if output != testCase.expected {
				// output was wrong
				t.Fatalf(`
Expected:
	%v
Received:
	%v
`, testCase.expected, output)

			}
		}

	}
}
