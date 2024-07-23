package models

import "time"

type Karyawan struct {
	ID        int       `json:"id"`
	Name      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
}
