package models

import "time"

type Category struct {
	ID        uint
	Name      string
	SortOrder int
	CreatedAt time.Time
	UpdatedAt time.Time
}
