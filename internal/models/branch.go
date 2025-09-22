package models

import "time"

type Branch struct {
	ID        uint
	Code      string
	Name      string
	CompanyID uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Company Company `gorm:"foreignKey:CompanyID"`
}
