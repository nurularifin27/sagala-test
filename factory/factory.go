package factory

import (
	"sagala/internal/repository"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Factory struct {
	DB *gorm.DB

	CompanyRepo      *repository.CompanyRepository
	BranchRepo       *repository.BranchRepository
	BrandRepo        *repository.BrandRepository
	CategoryRepo     *repository.CategoryRepository
	ChannelRepo      *repository.ChannelRepository
	MenuRepo         *repository.MenuRepository
	MerchantRepo     *repository.MerchantRepository
	MerchantMenuRepo *repository.MerchantMenuRepository
}

func NewFactory() *Factory {
	dsn := viper.GetString("database.dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	return &Factory{
		DB: db,

		CompanyRepo:      repository.NewCompanyRepository(db),
		BranchRepo:       repository.NewBranchRepository(db),
		BrandRepo:        repository.NewBrandRepository(db),
		CategoryRepo:     repository.NewCategoryRepository(db),
		ChannelRepo:      repository.NewChannelRepository(db),
		MenuRepo:         repository.NewMenuRepository(db),
		MerchantRepo:     repository.NewMerchantRepository(db),
		MerchantMenuRepo: repository.NewMerchantMenuRepository(db),
	}
}
