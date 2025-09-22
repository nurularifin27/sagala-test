package repository

import (
	"sagala/internal/models"

	"gorm.io/gorm"
)

type MerchantMenuRepository struct {
	db *gorm.DB
}

func NewMerchantMenuRepository(db *gorm.DB) *MerchantMenuRepository {
	return &MerchantMenuRepository{db: db}
}

func (r *MerchantMenuRepository) Create(merchantMenu *models.MerchantMenu) error {
	return r.db.Create(merchantMenu).Error
}

func (r *MerchantMenuRepository) FindAll() ([]models.MerchantMenu, error) {
	var merchantMenus []models.MerchantMenu
	result := r.db.Preload("Merchant").
		Preload("Menu").
		Preload("Category").
		Order("sort_order asc").
		Find(&merchantMenus)
	return merchantMenus, result.Error
}

func (r *MerchantMenuRepository) FindByID(id uint) (models.MerchantMenu, error) {
	var merchantMenu models.MerchantMenu
	result := r.db.Preload("Merchant").
		Preload("Menu").
		Preload("Category").
		First(&merchantMenu, id)
	return merchantMenu, result.Error
}

func (r *MerchantMenuRepository) FindByMerchantID(merchantID uint) ([]models.MerchantMenu, error) {
	var merchantMenus []models.MerchantMenu
	result := r.db.Where("merchant_id = ?", merchantID).
		Preload("Menu").
		Preload("Category").
		Order("sort_order asc").
		Find(&merchantMenus)
	return merchantMenus, result.Error
}

func (r *MerchantMenuRepository) FindByMenuID(menuID uint) ([]models.MerchantMenu, error) {
	var merchantMenus []models.MerchantMenu
	result := r.db.Where("menu_id = ?", menuID).
		Preload("Menu").
		Preload("Category").
		Order("sort_order asc").
		Find(&merchantMenus)
	return merchantMenus, result.Error
}

func (r *MerchantMenuRepository) FindByIDAndMerchantID(id, merchantID uint) (models.MerchantMenu, error) {
	var merchantMenu models.MerchantMenu
	result := r.db.Where("id = ? AND merchant_id = ?", id, merchantID).
		Preload("Menu").
		Preload("Category").
		First(&merchantMenu)
	return merchantMenu, result.Error
}

func (r *MerchantMenuRepository) Update(merchantMenu *models.MerchantMenu) error {
	return r.db.Save(merchantMenu).Error
}

func (r *MerchantMenuRepository) UpdatePrice(merchantID, menuID uint, price, discount float64) error {
	return r.db.Model(&models.MerchantMenu{}).
		Where("merchant_id = ? AND menu_id = ?", merchantID, menuID).
		Updates(map[string]interface{}{"price": price, "discount": discount}).Error
}

func (r *MerchantMenuRepository) BulkUpdatePriceByMenuID(menuID uint, price, discount float64) error {
	return r.db.Model(&models.MerchantMenu{}).
		Where("menu_id = ?", menuID).
		Updates(map[string]interface{}{"price": price, "discount": discount}).Error
}

func (r *MerchantMenuRepository) Delete(id uint) error {
	return r.db.Delete(&models.MerchantMenu{}, id).Error
}
