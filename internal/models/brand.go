package models

import "time"

type Brand struct {
	ID        uint
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
