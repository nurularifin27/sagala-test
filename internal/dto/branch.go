package dto

import "sagala/internal/models"

type BranchFilterRequest struct {
	CompanyID uint `query:"company_id"`
}

type BranchRequest struct {
	Code      string `json:"code" binding:"required"`
	Name      string `json:"name" binding:"required"`
	CompanyID uint   `json:"company_id" binding:"required"`
}

type BranchResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (br *BranchResponse) BuildListBranchResponse(branches []models.Branch) []BranchResponse {
	var responses []BranchResponse
	for _, branch := range branches {
		responses = append(responses, BranchResponse{
			ID:   branch.ID,
			Code: branch.Code,
			Name: branch.Name,
		})
	}
	return responses
}

func (br *BranchResponse) BuildBranchResponse(branch models.Branch) *BranchResponse {
	return &BranchResponse{
		ID:   branch.ID,
		Code: branch.Code,
		Name: branch.Name,
	}
}
