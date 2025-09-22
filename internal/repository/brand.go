package repository

import (
	"sagala/internal/models"

	"gorm.io/gorm"
)

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{db: db}
}

func (r *BrandRepository) Create(brand *models.Brand) error {
	return r.db.Create(brand).Error
}

func (r *BrandRepository) FindAll() ([]models.Brand, error) {
	var companies []models.Brand
	result := r.db.Find(&companies)
	return companies, result.Error
}

func (r *BrandRepository) FindByID(id uint) (models.Brand, error) {
	var brand models.Brand
	result := r.db.First(&brand, id)
	return brand, result.Error
}

func (r *BrandRepository) FindByCode(code string) (models.Brand, error) {
	var brand models.Brand
	result := r.db.Where("code = ?", code).First(&brand)
	return brand, result.Error
}

func (r *BrandRepository) Update(brand *models.Brand) error {
	return r.db.Save(brand).Error
}

func (r *BrandRepository) Delete(id uint) error {
	return r.db.Delete(&models.Brand{}, id).Error
}
