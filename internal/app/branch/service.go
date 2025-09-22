package branch

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"

	"gorm.io/gorm"
)

type BranchService struct {
	branchRepo *repository.BranchRepository
}

func NewBranchService(factory *factory.Factory) *BranchService {
	return &BranchService{
		branchRepo: factory.BranchRepo,
	}
}

func (s *BranchService) Create(request *dto.BranchRequest) error {
	branch := models.Branch{
		Code:      request.Code,
		Name:      request.Name,
		CompanyID: request.CompanyID,
	}

	if err := s.branchRepo.Create(&branch); err != nil {
		return err
	}

	return nil
}

func (s *BranchService) Update(id uint, request *dto.BranchRequest) error {
	branch, err := s.branchRepo.FindByID(id)
	if err != nil {
		return err
	}

	branch.Code = request.Code
	branch.Name = request.Name
	branch.CompanyID = request.CompanyID

	return s.branchRepo.Update(&branch)
}

func (s *BranchService) Delete(id uint) error {
	return s.branchRepo.Delete(id)
}

func (s *BranchService) Index(filter *dto.BranchFilterRequest) ([]dto.BranchResponse, error) {
	var branches []models.Branch
	var err error
	var resp dto.BranchResponse
	if filter.CompanyID > 0 {
		branches, err = s.branchRepo.FindByCompanyID(filter.CompanyID)
	} else {
		branches, err = s.branchRepo.FindAll()
	}
	if err != nil {
		return nil, err
	}

	return resp.BuildListBranchResponse(branches), nil
}

func (s *BranchService) Show(id uint) (*dto.BranchResponse, error) {
	var resp dto.BranchResponse
	branch, err := s.branchRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return resp.BuildBranchResponse(branch), nil
}
