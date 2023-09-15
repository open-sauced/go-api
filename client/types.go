package client

import (
	"fmt"
	"strings"
	"time"
)

// DateTime is a bridge between a swagger doc "date-time" and a Golang time.Time type
type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		dt.Time = time.Time{}
		return nil
	}
	parsedTime, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}
	dt.Time = parsedTime
	return nil
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", dt.Time.Format(time.DateOnly))), nil
}
