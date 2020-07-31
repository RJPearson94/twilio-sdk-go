package utils

import (
	"encoding/json"
	"time"
)

// RFC2822Time allows RFC2822 returned by the Twilio APIs to marshalled/ unmarshalled
type RFC2822Time struct {
	*time.Time
}

// RFC2822 is the time format used by Twilio's unified API
const RFC2822 = "Mon, 2 Jan 2006 15:04:05 -0700"

func (value *RFC2822Time) UnmarshalJSON(b []byte) error {
	t := new(string)
	if err := json.Unmarshal(b, t); err != nil {
		return err
	}

	if t == nil || *t == "null" || *t == "" {
		*value = RFC2822Time{Time: nil}
		return nil
	}

	parsedTime, err := time.Parse(RFC2822, *t)
	if err != nil {
		return err
	}

	*value = RFC2822Time{Time: &parsedTime}
	return nil
}

func (value *RFC2822Time) MarshalJSON() ([]byte, error) {
	byteArray, err := json.Marshal(value.Time)
	if err != nil {
		return []byte{}, err
	}
	return byteArray, nil
}
