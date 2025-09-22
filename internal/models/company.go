package models

import "time"

type Company struct {
	ID        uint
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
