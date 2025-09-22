package models

import "time"

type MerchantMenu struct {
	ID         uint
	MerchantID uint
	MenuID     uint
	CategoryID uint
	SortOrder  int
	Price      float64
	Discount   float64
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Merchant Merchant `gorm:"foreignKey:MerchantID"`
	Menu     Menu     `gorm:"foreignKey:MenuID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}
