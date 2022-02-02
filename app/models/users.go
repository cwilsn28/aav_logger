package models

import "time"

type User struct {
	ID        int64     `json:"user_id"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	Role      string    `json:"role"`
	Created   time.Time `json:"created"`
	LastLogin time.Time `json:"last_login"`
	Status    int64     `json:"status"`
}
