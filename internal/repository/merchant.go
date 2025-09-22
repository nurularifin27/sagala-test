package repository

import (
	"sagala/internal/models"

	"gorm.io/gorm"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

func (r *MerchantRepository) Create(merchant *models.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *MerchantRepository) FindAll() ([]models.Merchant, error) {
	var merchants []models.Merchant
	result := r.db.Preload("Branch").Preload("Brand").Preload("Channel").Find(&merchants)
	return merchants, result.Error
}

func (r *MerchantRepository) FindByID(id uint) (models.Merchant, error) {
	var merchant models.Merchant
	result := r.db.Preload("Branch").Preload("Brand").Preload("Channel").First(&merchant, id)
	return merchant, result.Error
}

func (r *MerchantRepository) FindByBranchID(branchID uint) ([]models.Merchant, error) {
	var merchants []models.Merchant
	result := r.db.Where("branch_id = ?", branchID).Preload("Branch").Preload("Brand").Preload("Channel").Find(&merchants)
	return merchants, result.Error
}

func (r *MerchantRepository) FindByBranchIDAndBrandIDAndChannelID(branchID, brandID, channelID uint) (models.Merchant, error) {
	var merchant models.Merchant
	result := r.db.Where("branch_id = ? AND brand_id = ? AND channel_id = ?", branchID, brandID, channelID).
		Preload("Branch").
		Preload("Brand").
		Preload("Channel").
		First(&merchant)
	return merchant, result.Error
}

func (r *MerchantRepository) Update(merchant *models.Merchant) error {
	return r.db.Save(merchant).Error
}

func (r *MerchantRepository) Delete(id uint) error {
	return r.db.Delete(&models.Merchant{}, id).Error
}
