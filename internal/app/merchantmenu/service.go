package merchantmenu

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"
)

type MerchantMenuService struct {
	merchantMenuRepo *repository.MerchantMenuRepository
	menuRepo         *repository.MenuRepository
}

func NewMerchantMenuService(factory *factory.Factory) *MerchantMenuService {
	return &MerchantMenuService{
		merchantMenuRepo: factory.MerchantMenuRepo,
		menuRepo:         factory.MenuRepo,
	}
}

func (s *MerchantMenuService) Create(request *dto.MerchantMenuRequest) error {
	merchantMenu := models.MerchantMenu{
		MerchantID: request.MerchantID,
		MenuID:     request.MenuID,
		CategoryID: request.CategoryID,
		SortOrder:  request.SortOrder,
		Price:      request.Price,
		Discount:   request.Discount,
	}

	if err := s.merchantMenuRepo.Create(&merchantMenu); err != nil {
		return err
	}

	return nil
}

func (s *MerchantMenuService) Update(id uint, request *dto.MerchantMenuRequest) error {
	merchantMenu, err := s.merchantMenuRepo.FindByID(id)
	if err != nil {
		return err
	}

	merchantMenu.MenuID = request.MenuID
	merchantMenu.CategoryID = request.CategoryID
	merchantMenu.SortOrder = request.SortOrder
	merchantMenu.Price = request.Price
	merchantMenu.Discount = request.Discount

	return s.merchantMenuRepo.Update(&merchantMenu)
}

func (s *MerchantMenuService) Delete(id uint) error {
	return s.merchantMenuRepo.Delete(id)
}

func (s *MerchantMenuService) Index(filter *dto.MerchantMenuFilterRequest) ([]dto.MerchantMenuResponse, error) {
	var resp dto.MerchantMenuResponse
	merchantMenus, err := s.merchantMenuRepo.FindByMerchantID(filter.MerchantID)
	if err != nil {
		return nil, err
	}

	return resp.BuildListMerchantMenuResponse(merchantMenus), nil
}

func (s *MerchantMenuService) Show(id uint) (*dto.MerchantMenuResponse, error) {
	var resp dto.MerchantMenuResponse
	merchantMenu, err := s.merchantMenuRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return resp.BuildMerchantMenuResponse(merchantMenu), nil
}

func (s *MerchantMenuService) UpdatePrice(request *dto.MerchantMenuPriceRequest) error {
	if err := s.merchantMenuRepo.BulkUpdatePriceByMenuID(request.MenuID, request.Price, request.Discount); err != nil {
		return err
	}

	return nil
}

func (s *MerchantMenuService) MenuBranch(branchId uint) ([]dto.MenuResponse, error) {
	var resp dto.MenuResponse
	menus, err := s.menuRepo.FindByBranchID(branchId)
	if err != nil {
		return nil, err
	}
	return resp.BuildListMenuResponse(menus), nil
}

func (s *MerchantMenuService) MenuMerchant(merchantId uint) ([]dto.MenuResponse, error) {
	var resp dto.MenuResponse
	merchantMenus, err := s.merchantMenuRepo.FindByMerchantID(merchantId)
	if err != nil {
		return nil, err
	}

	menus := make([]models.Menu, 0)

	for _, merchantMenu := range merchantMenus {
		menus = append(menus, merchantMenu.Menu)
	}

	return resp.BuildListMenuResponse(menus), nil
}
