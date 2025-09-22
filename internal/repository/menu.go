package repository

import (
	"sagala/internal/models"

	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) Create(menu *models.Menu) error {
	return r.db.Create(menu).Error
}

func (r *MenuRepository) FindAll() ([]models.Menu, error) {
	var menus []models.Menu
	result := r.db.Find(&menus)
	return menus, result.Error
}

func (r *MenuRepository) FindByID(id uint) (models.Menu, error) {
	var menu models.Menu
	result := r.db.First(&menu, id)
	return menu, result.Error
}

func (r *MenuRepository) Update(menu *models.Menu) error {
	return r.db.Save(menu).Error
}

func (r *MenuRepository) Delete(id uint) error {
	return r.db.Delete(&models.Menu{}, id).Error
}

func (r *MenuRepository) FindByBranchID(branchID uint) ([]models.Menu, error) {
	var menus []models.Menu

	err := r.db.
		Model(&models.Menu{}).
		Select("DISTINCT menus.*").
		Joins("JOIN merchant_menus mm ON mm.menu_id = menus.id").
		Joins("JOIN merchants m ON m.id = mm.merchant_id").
		Where("m.branch_id = ?", branchID).
		Find(&menus).Error

	if err != nil {
		return nil, err
	}
	return menus, nil
}
