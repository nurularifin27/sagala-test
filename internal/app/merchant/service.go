package merchant

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"

	"gorm.io/gorm"
)

type MerchantService struct {
	merchantRepo *repository.MerchantRepository
}

func NewMerchantService(factory *factory.Factory) *MerchantService {
	return &MerchantService{
		merchantRepo: factory.MerchantRepo,
	}
}

func (s *MerchantService) Create(request *dto.MerchantRequest) error {
	merchant := models.Merchant{
		BranchID:  request.BranchID,
		BrandID:   request.BrandID,
		ChannelID: request.ChannelID,
	}

	if err := s.merchantRepo.Create(&merchant); err != nil {
		return err
	}

	return nil
}

func (s *MerchantService) Update(id uint, request *dto.MerchantRequest) error {
	merchant, err := s.merchantRepo.FindByID(id)
	if err != nil {
		return err
	}

	merchant.BranchID = request.BranchID
	merchant.BrandID = request.BrandID
	merchant.ChannelID = request.ChannelID

	return s.merchantRepo.Update(&merchant)
}

func (s *MerchantService) Delete(id uint) error {
	return s.merchantRepo.Delete(id)
}

func (s *MerchantService) Index() ([]dto.MerchantResponse, error) {
	var resp dto.MerchantResponse
	merchants, err := s.merchantRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return resp.BuildListMerchantResponse(merchants), nil
}

func (s *MerchantService) Show(id uint) (*dto.MerchantResponse, error) {
	var resp dto.MerchantResponse
	merchant, err := s.merchantRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return resp.BuildMerchantResponse(merchant), nil
}
