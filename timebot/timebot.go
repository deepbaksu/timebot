// Package timebot is a service package providing
// helper functions neeeded for time conversions
package timebot

import (
	"errors"
	"fmt"
	s "strings"
	"time"
)

func CheckDaylightSavingZone(text string) (daylightSavingZone bool) {

	tzDb := map[string]string{
		"PST": "America/Los_Angeles",
		"PDT": "America/Los_Angeles",
	}

	tzText := ""
	for k, v := range tzDb {
		if s.Contains(text, k) {
			tzText = v
		}
	}
	if tzText == "" {
		return true
	}

	const longForm2 = "2006-01-02 15:04 MST"
	loc, _ := time.LoadLocation(tzText)

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

//nolint:gochecknoglobals
var (
	koTz, _ = time.LoadLocation("Asia/Seoul")
	caTz, _ = time.LoadLocation("America/Los_Angeles")
)

// ParseAndFlipTz returns a datetime string but in other TZ
//
// The following conversions are supported as of now:
//
// 1. KST <-> PST/PDT
func ParseAndFlipTz(text string) (string, error) {
	tzFlipDb := map[string]*time.Location{
		"PST": koTz,
		"PDT": koTz,
		"KST": caTz,
	}

	for tzText, tz := range tzFlipDb {
		if s.Contains(text, tzText) {
			utc, ok := ParseTime(text)

			if !ok {
				return utc.String(), errors.New("fail to ParseTime()")
			}

			tString := utc.In(tz).Format("2006-01-02 15:04 MST")
			return tString, nil
		}
	}

	return "", fmt.Errorf("%v does not contain PST/PDT/KST", text)
}
