package repository

import (
	"sagala/internal/models"

	"gorm.io/gorm"
)

type ChannelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	return &ChannelRepository{db: db}
}

func (r *ChannelRepository) Create(channel *models.Channel) error {
	return r.db.Create(channel).Error
}

func (r *ChannelRepository) FindAll() ([]models.Channel, error) {
	var companies []models.Channel
	result := r.db.Find(&companies)
	return companies, result.Error
}

func (r *ChannelRepository) FindByID(id uint) (models.Channel, error) {
	var channel models.Channel
	result := r.db.First(&channel, id)
	return channel, result.Error
}

func (r *ChannelRepository) FindByCode(code string) (models.Channel, error) {
	var channel models.Channel
	result := r.db.Where("code = ?", code).First(&channel)
	return channel, result.Error
}

func (r *ChannelRepository) Update(channel *models.Channel) error {
	return r.db.Save(channel).Error
}

func (r *ChannelRepository) Delete(id uint) error {
	return r.db.Delete(&models.Channel{}, id).Error
}
