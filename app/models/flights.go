package models

import (
	"strings"
	"time"
)

type Flight struct {
	ID         int64     `json:"id"`
	Robot      string    `json:"robot"`
	Generation int64     `json:"generation"`
	Start      time.Time `json:"start"`
	Stop       time.Time `json:"stop"`
	Lat        float64   `json:"lat"`
	Lon        float64   `json:"lon"`
}

func (f Flight) IsValid() (bool, string) {
	var msgs = make([]string, 0)
	var valid = true

	if f.Robot == "" {
		valid = false
		msgs = append(msgs, "missing robot name")
	}
	if f.Generation == 0 {
		valid = false
		msgs = append(msgs, "missing robot generation")
	}
	if f.Start.IsZero() {
		valid = false
		msgs = append(msgs, "missing or invalid start time")
	}
	if f.Stop.IsZero() {
		valid = false
		msgs = append(msgs, "missing or invalid stop time")
	}
	if f.Lat == 0 {
		valid = false
		msgs = append(msgs, "missing or invalid latitude")
	}
	if f.Lat == 0 {
		valid = false
		msgs = append(msgs, "missing or invalid longitude")
	}
	return valid, strings.Join(msgs, ", ")
}
