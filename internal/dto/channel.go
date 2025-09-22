package dto

import "sagala/internal/models"

type ChannelRequest struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ChannelResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (cr *ChannelResponse) BuildListChannelResponse(channels []models.Channel) []ChannelResponse {
	var responses []ChannelResponse
	for _, channel := range channels {
		responses = append(responses, ChannelResponse{
			ID:   channel.ID,
			Code: channel.Code,
			Name: channel.Name,
		})
	}
	return responses
}

func (cr *ChannelResponse) BuildChannelResponse(channel models.Channel) *ChannelResponse {
	return &ChannelResponse{
		ID:   channel.ID,
		Code: channel.Code,
		Name: channel.Name,
	}
}
