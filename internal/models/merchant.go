package models

import "time"

type Merchant struct {
	ID        uint
	BranchID  uint
	BrandID   uint
	ChannelID uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Branch  Branch  `gorm:"foreignKey:BranchID"`
	Brand   Brand   `gorm:"foreignKey:BrandID"`
	Channel Channel `gorm:"foreignKey:ChannelID"`
}
