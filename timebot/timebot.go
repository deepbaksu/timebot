// Package timebot is a service package providing
// helper functions neeeded for time conversions
package timebot

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	s "strings"
	"time"
)

//nolint:gochecknoglobals
var (
	koTz, _ = time.LoadLocation("Asia/Seoul")
	caTz, _ = time.LoadLocation("America/Los_Angeles")
)

//nolint:gochecknoglobals
var pstPdt = regexp.MustCompile("PST|PDT")

// ParseTime parses datetime string to time.Time object
// It also returns boolean whether parsing was successful
func ParseTime(text string) (time.Time, bool) {

	tzToTimeGap := map[string]string{
		"KST": "+0900",
	}

	log.Println("ParseTime", text)
	if s.Contains(text, "PST") || s.Contains(text, "PDT") {
		// (1) remove PST/PDT
		// (2) parse using timezone object `America/Los_Angeles`
		const pstpdtForm = "2006-01-02 15:04"
		text = s.Trim(pstPdt.ReplaceAllString(text, ""), " ")
		t, err := time.ParseInLocation(pstpdtForm, text, caTz)

		if err != nil {
			log.Fatalln("Failed to parse PST/PDT: ", err)
			return t, false
		}
		return t.UTC(), true
	}

	const longForm = "2006-01-02 15:04 -0700"

	for key, value := range tzToTimeGap {
		text = s.Replace(text, key, value, 1)
	}

	t, err := time.Parse(longForm, text)

	if err != nil {
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

			tString := utc.In(tz).Format("2006-01-02 15:04 MST (Mon)")
			return tString, nil
		}
	}

	return "", fmt.Errorf("%v does not contain PST/PDT/KST", text)
}

// Regex for 2006-01-02 15:04 MST
// nolint:gochecknoglobals
var datetimeRegex = regexp.MustCompile(`\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}\s[A-z]{3}`)

// ExtractDateTime extracts DateTime from 'text'
func ExtractDateTime(text string) (string, error) {
	find := datetimeRegex.FindString(text)

	if find != "" {
		return find, nil
	}

	return "", errors.New("text does not contain valid date time format (e.g., 2006-01-02 15:04 MST)")
}
