package repository

import (
	"sagala/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Order("sort_order asc").Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) FindByID(id uint) (models.Category, error) {
	var category models.Category
	result := r.db.First(&category, id)
	return category, result.Error
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
