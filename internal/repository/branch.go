package repository

import (
	"sagala/internal/models"

	"gorm.io/gorm"
)

type BranchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) *BranchRepository {
	return &BranchRepository{db: db}
}

func (r *BranchRepository) Create(branch *models.Branch) error {
	return r.db.Create(branch).Error
}

func (r *BranchRepository) FindAll() ([]models.Branch, error) {
	var branches []models.Branch
	result := r.db.Preload("Company").Find(&branches)
	return branches, result.Error
}

func (r *BranchRepository) FindByID(id uint) (models.Branch, error) {
	var branch models.Branch
	result := r.db.Preload("Company").First(&branch, id)
	return branch, result.Error
}

func (r *BranchRepository) FindByCompanyID(companyID uint) ([]models.Branch, error) {
	var branches []models.Branch
	result := r.db.Where("company_id = ?", companyID).Preload("Company").Find(&branches)
	return branches, result.Error
}

func (r *BranchRepository) Update(branch *models.Branch) error {
	return r.db.Save(branch).Error
}

func (r *BranchRepository) Delete(id uint) error {
	return r.db.Delete(&models.Branch{}, id).Error
}
