package dto

import "sagala/internal/models"

type MerchantMenuRequest struct {
	MerchantID uint    `json:"merchant_id" binding:"required"`
	MenuID     uint    `json:"menu_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
	SortOrder  int     `json:"sort_order"`
	Price      float64 `json:"price" binding:"required,gt=0"`
	Discount   float64 `json:"discount"`
}

type MerchantMenuFilterRequest struct {
	MerchantID uint `query:"merchant_id"`
}

type MerchantMenuResponse struct {
	ID         uint              `json:"id"`
	Menu       *MenuResponse     `json:"menu"`
	Category   *CategoryResponse `json:"category"`
	SortOrder  int               `json:"sort_order"`
	Price      float64           `json:"price"`
	Discount   float64           `json:"discount"`
	FinalPrice float64           `json:"final_price"`
}

func (mmr *MerchantMenuResponse) BuildMerchantMenuResponse(merchantMenu models.MerchantMenu) *MerchantMenuResponse {
	return &MerchantMenuResponse{
		ID: merchantMenu.ID,
		Menu: &MenuResponse{
			ID:          merchantMenu.Menu.ID,
			Name:        merchantMenu.Menu.Name,
			ImageURL:    merchantMenu.Menu.ImageURL,
			Description: merchantMenu.Menu.Description,
		},
		Category: &CategoryResponse{
			ID:        merchantMenu.Category.ID,
			Name:      merchantMenu.Category.Name,
			SortOrder: merchantMenu.Category.SortOrder,
		},
		SortOrder:  merchantMenu.SortOrder,
		Price:      merchantMenu.Price,
		Discount:   merchantMenu.Discount,
		FinalPrice: merchantMenu.Price - merchantMenu.Discount,
	}
}

func (mmr *MerchantMenuResponse) BuildListMerchantMenuResponse(merchantMenus []models.MerchantMenu) []MerchantMenuResponse {
	var responses []MerchantMenuResponse
	for _, merchantMenu := range merchantMenus {
		responses = append(responses, *mmr.BuildMerchantMenuResponse(merchantMenu))
	}
	return responses
}

type MerchantMenuPriceRequest struct {
	MenuID   uint    `json:"menu_id" binding:"required"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Discount float64 `json:"discount"`
}

type MenusByMerchantResponse struct {
	Merchant MerchantResponse     `json:"merchant"`
	Menu     MenuMerchantResponse `json:"menu"`
}

type MenuMerchantResponse struct {
	Categories []CategoryMenuResponse `json:"categories"`
}
type MenusByBranchResponse struct {
	Branch    BranchResponse            `json:"branch"`
	Merchants []MenusByMerchantResponse `json:"merchants"`
}
