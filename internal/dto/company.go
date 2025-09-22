package dto

import "sagala/internal/models"

type CompanyRequest struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CompanyResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (cr *CompanyResponse) BuildListCompanyResponse(companies []models.Company) []CompanyResponse {
	var responses []CompanyResponse
	for _, company := range companies {
		responses = append(responses, CompanyResponse{
			ID:   company.ID,
			Code: company.Code,
			Name: company.Name,
		})
	}
	return responses
}

func (cr *CompanyResponse) BuildCompanyResponse(company models.Company) *CompanyResponse {
	return &CompanyResponse{
		ID:   company.ID,
		Code: company.Code,
		Name: company.Name,
	}
}
