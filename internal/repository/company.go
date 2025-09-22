package repository

import (
	"sagala/internal/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepository) FindAll() ([]models.Company, error) {
	var companies []models.Company
	result := r.db.Find(&companies)
	return companies, result.Error
}

func (r *CompanyRepository) FindByID(id uint) (models.Company, error) {
	var company models.Company
	result := r.db.First(&company, id)
	return company, result.Error
}

func (r *CompanyRepository) FindByCode(code string) (models.Company, error) {
	var company models.Company
	result := r.db.Where("code = ?", code).First(&company)
	return company, result.Error
}

func (r *CompanyRepository) Update(company *models.Company) error {
	return r.db.Save(company).Error
}

func (r *CompanyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Company{}, id).Error
}
