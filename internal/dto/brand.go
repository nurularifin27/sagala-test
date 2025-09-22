package dto

import "sagala/internal/models"

type BrandRequest struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type BrandResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (br *BrandResponse) BuildListBrandResponse(brands []models.Brand) []BrandResponse {
	var responses []BrandResponse
	for _, brand := range brands {
		responses = append(responses, BrandResponse{
			ID:   brand.ID,
			Code: brand.Code,
			Name: brand.Name,
		})
	}
	return responses
}

func (br *BrandResponse) BuildBrandResponse(brand models.Brand) *BrandResponse {
	return &BrandResponse{
		ID:   brand.ID,
		Code: brand.Code,
		Name: brand.Name,
	}
}

type BrandMenuResponse struct {
	Brand      BrandResponse          `json:"brand"`
	Categories []CategoryMenuResponse `json:"categories"`
}
