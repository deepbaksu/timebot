// Package timebot is a service package providing
// helper functions neeeded for time conversions
package timebot

import (
	"fmt"
	s "strings"
	"time"
)

func CheckDaylightSavingZone(text string) (daylightSavingZone bool) {

	tzDb := map[string]string{
		"PST": "America/Los_Angeles",
		"PDT": "America/Los_Angeles",
	}

	var tzText string = ""

	for k, v := range tzDb {
		if s.Contains(text, k) {
			tzText = v
		}
	}
	if tzText == "" {
		return true
	}

	const longForm2 = "2006-01-02 15:04 MST"
	loc, err := time.LoadLocation(tzText)
	if err != nil {
		fmt.Println("wrong tzDb")
	}

	t, _ := time.ParseInLocation(longForm2, text, loc)
	tString := t.Format("2006-01-02 15:04 MST")
	return text == tString
}

func ParseTime(text string) (time.Time, bool) {

	tzToTimeGap := map[string]string{
		"PST": "-0800",
		"PDT": "-0700",
		"KST": "+0900",
	}

	passCheck := CheckDaylightSavingZone(text)

	const longForm = "2006-01-02 15:04 -0700"

	for key, value := range tzToTimeGap {
		text = s.Replace(text, key, value, 1)
	}

	t, err := time.Parse(longForm, text)
	if err != nil || !passCheck {
		return t, false
	}
	return t.UTC(), true
}

// ParseAndFlipTz returns a datetime string but in other TZ
//
// The following conversions are supported as of now:
//
// 1. KST <-> PST/PDT
func ParseAndFlipTz(text string) (string, error) {
	return "", nil
}
