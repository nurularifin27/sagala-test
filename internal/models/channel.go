package models

import "time"

type Channel struct {
	ID        uint
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
