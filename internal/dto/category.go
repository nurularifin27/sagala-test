package dto

import "sagala/internal/models"

type CategoryRequest struct {
	Name      string `json:"name" binding:"required"`
	SortOrder int    `json:"sort_order" binding:"required"`
}

type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

func (cr *CategoryResponse) BuildListCategoryResponse(categories []models.Category) []CategoryResponse {
	var responses []CategoryResponse
	for _, category := range categories {
		responses = append(responses, CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			SortOrder: category.SortOrder,
		})
	}
	return responses
}

func (cr *CategoryResponse) BuildCategoryResponse(category models.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		SortOrder: category.SortOrder,
	}
}

type CategoryMenuResponse struct {
	ID        uint                    `json:"id"`
	Name      string                  `json:"name"`
	SortOrder int                     `json:"sort_order"`
	Items     []MenuWithPriceResponse `json:"items"`
}
