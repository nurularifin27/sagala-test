package brand

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"

	"gorm.io/gorm"
)

type BrandService struct {
	brandRepo *repository.BrandRepository
}

func NewBrandService(factory *factory.Factory) *BrandService {
	return &BrandService{
		brandRepo: factory.BrandRepo,
	}
}

func (s *BrandService) Create(request *dto.BrandRequest) error {
	brand := models.Brand{
		Code: request.Code,
		Name: request.Name,
	}

	if err := s.brandRepo.Create(&brand); err != nil {
		return err
	}

	return nil
}

func (s *BrandService) Update(id uint, request *dto.BrandRequest) error {
	brand, err := s.brandRepo.FindByID(id)
	if err != nil {
		return err
	}

	brand.Code = request.Code
	brand.Name = request.Name

	return s.brandRepo.Update(&brand)
}

func (s *BrandService) Delete(id uint) error {
	return s.brandRepo.Delete(id)
}

func (s *BrandService) Index() ([]dto.BrandResponse, error) {
	var resp dto.BrandResponse
	brands, err := s.brandRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return resp.BuildListBrandResponse(brands), nil
}

func (s *BrandService) Show(id uint) (*dto.BrandResponse, error) {
	var resp dto.BrandResponse
	brand, err := s.brandRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return resp.BuildBrandResponse(brand), nil
}
