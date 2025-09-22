package dto

import "sagala/internal/models"

type MenuRequest struct {
	Name        string  `json:"name" binding:"required"`
	ImageURL    *string `json:"image_url"`
	Description *string `json:"description"`
}

type MenuResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	ImageURL    *string `json:"image_url"`
	Description *string `json:"description"`
}

type MenuWithPriceResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	ImageURL    *string `json:"image_url"`
	Description *string `json:"description"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	FinalPrice  float64 `json:"final_price"`
	SortOrder   int     `json:"sort_order"`
}

func (mr *MenuResponse) BuildListMenuResponse(companies []models.Menu) []MenuResponse {
	var responses []MenuResponse
	for _, company := range companies {
		responses = append(responses, MenuResponse{
			ID:          company.ID,
			Name:        company.Name,
			ImageURL:    company.ImageURL,
			Description: company.Description,
		})
	}
	return responses
}

func (mr *MenuResponse) BuildMenuResponse(company models.Menu) *MenuResponse {
	return &MenuResponse{
		ID:          company.ID,
		Name:        company.Name,
		ImageURL:    company.ImageURL,
		Description: company.Description,
	}
}
