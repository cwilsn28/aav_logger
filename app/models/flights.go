package models

import "time"

type Flight struct {
	ID         int64     `json:"id"`
	Robot      string    `json:"robot"`
	Generation int64     `json:"generation"`
	Start      time.Time `json:"start"`
	Stop       time.Time `json:"stop"`
	Lat        float64   `json:"lat"`
	Lon        float64   `json:"lon"`
}
