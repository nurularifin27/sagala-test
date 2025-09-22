package models

import "time"

type Menu struct {
	ID          uint
	Name        string
	ImageURL    *string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	MerchantMenus []MerchantMenu `gorm:"foreignKey:MenuID"`
}
