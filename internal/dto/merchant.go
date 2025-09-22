package dto

import "sagala/internal/models"

type MerchantRequest struct {
	BranchID  uint `json:"branch_id" binding:"required"`
	BrandID   uint `json:"brand_id" binding:"required"`
	ChannelID uint `json:"channel_id" binding:"required"`
}

type MerchantResponse struct {
	ID      uint             `json:"id"`
	Branch  *BranchResponse  `json:"branch"`
	Brand   *BrandResponse   `json:"brand"`
	Channel *ChannelResponse `json:"channel"`
}

func (mr *MerchantResponse) BuildListMerchantResponse(merchants []models.Merchant) []MerchantResponse {
	var responses []MerchantResponse
	for _, merchant := range merchants {
		responses = append(responses, MerchantResponse{
			ID: merchant.ID,
			Branch: &BranchResponse{
				ID:   merchant.Branch.ID,
				Name: merchant.Branch.Name,
				Code: merchant.Branch.Code,
			},
			Brand: &BrandResponse{
				ID:   merchant.Brand.ID,
				Code: merchant.Brand.Code,
				Name: merchant.Brand.Name,
			},
			Channel: &ChannelResponse{
				ID:   merchant.Channel.ID,
				Code: merchant.Channel.Code,
				Name: merchant.Channel.Name,
			},
		})
	}
	return responses
}

func (mr *MerchantResponse) BuildMerchantResponse(merchant models.Merchant) *MerchantResponse {
	return &MerchantResponse{
		ID: merchant.ID,
		Branch: &BranchResponse{
			ID:   merchant.Branch.ID,
			Name: merchant.Branch.Name,
			Code: merchant.Branch.Code,
		},
		Brand: &BrandResponse{
			ID:   merchant.Brand.ID,
			Code: merchant.Brand.Code,
			Name: merchant.Brand.Name,
		},
		Channel: &ChannelResponse{
			ID:   merchant.Channel.ID,
			Code: merchant.Channel.Code,
			Name: merchant.Channel.Name,
		},
	}
}
