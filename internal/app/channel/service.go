package channel

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"

	"gorm.io/gorm"
)

type ChannelService struct {
	channelRepo *repository.ChannelRepository
}

func NewChannelService(factory *factory.Factory) *ChannelService {
	return &ChannelService{
		channelRepo: factory.ChannelRepo,
	}
}

func (s *ChannelService) Create(request *dto.ChannelRequest) error {
	channel := models.Channel{
		Code: request.Code,
		Name: request.Name,
	}

	if err := s.channelRepo.Create(&channel); err != nil {
		return err
	}

	return nil
}

func (s *ChannelService) Update(id uint, request *dto.ChannelRequest) error {
	channel, err := s.channelRepo.FindByID(id)
	if err != nil {
		return err
	}

	channel.Code = request.Code
	channel.Name = request.Name

	return s.channelRepo.Update(&channel)
}

func (s *ChannelService) Delete(id uint) error {
	return s.channelRepo.Delete(id)
}

func (s *ChannelService) Index() ([]dto.ChannelResponse, error) {
	var resp dto.ChannelResponse
	channels, err := s.channelRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return resp.BuildListChannelResponse(channels), nil
}

func (s *ChannelService) Show(id uint) (*dto.ChannelResponse, error) {
	var resp dto.ChannelResponse
	channel, err := s.channelRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return resp.BuildChannelResponse(channel), nil
}
