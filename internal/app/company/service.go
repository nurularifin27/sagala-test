package company

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"

	"gorm.io/gorm"
)

type CompanyService struct {
	companyRepo *repository.CompanyRepository
}

func NewCompanyService(factory *factory.Factory) *CompanyService {
	return &CompanyService{
		companyRepo: factory.CompanyRepo,
	}
}

func (s *CompanyService) Create(request *dto.CompanyRequest) error {
	company := models.Company{
		Code: request.Code,
		Name: request.Name,
	}

	if err := s.companyRepo.Create(&company); err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) Update(id uint, request *dto.CompanyRequest) error {
	company, err := s.companyRepo.FindByID(id)
	if err != nil {
		return err
	}

	company.Code = request.Code
	company.Name = request.Name

	return s.companyRepo.Update(&company)
}

func (s *CompanyService) Delete(id uint) error {
	return s.companyRepo.Delete(id)
}

func (s *CompanyService) Index() ([]dto.CompanyResponse, error) {
	var resp dto.CompanyResponse
	companies, err := s.companyRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return resp.BuildListCompanyResponse(companies), nil
}

func (s *CompanyService) Show(id uint) (*dto.CompanyResponse, error) {
	var resp dto.CompanyResponse
	company, err := s.companyRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return resp.BuildCompanyResponse(company), nil
}
