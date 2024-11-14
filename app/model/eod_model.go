package model

import "time"

type Eod struct {
	ID        *int       `json:"id"`
	Name      *string    `json:"nama"`
	CreatedAt *time.Time `json:"created_at"`
}
