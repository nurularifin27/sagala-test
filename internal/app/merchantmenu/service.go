package merchantmenu

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"
	"sort"
)

type MerchantMenuService struct {
	merchantMenuRepo *repository.MerchantMenuRepository
	menuRepo         *repository.MenuRepository
	merchantRepo     *repository.MerchantRepository
	branchRepo       *repository.BranchRepository
}

func NewMerchantMenuService(factory *factory.Factory) *MerchantMenuService {
	return &MerchantMenuService{
		merchantMenuRepo: factory.MerchantMenuRepo,
		menuRepo:         factory.MenuRepo,
		merchantRepo:     factory.MerchantRepo,
		branchRepo:       factory.BranchRepo,
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

func (s *MerchantMenuService) MenuBranch(branchId uint) (*dto.MenusByBranchResponse, error) {
	branch, err := s.branchRepo.FindByID(branchId)
	if err != nil {
		return nil, err
	}

	merchants, err := s.merchantRepo.FindByBranchID(branchId)
	if err != nil {
		return nil, err
	}

	menuMerchants := make([]dto.MenusByMerchantResponse, 0)
	for _, merchant := range merchants {
		menuMerchant, err := s.MenuMerchant(merchant.ID)
		if err != nil {
			return nil, err
		}
		menuMerchants = append(menuMerchants, *menuMerchant)
	}

	return &dto.MenusByBranchResponse{
		Branch: dto.BranchResponse{
			ID:   branch.ID,
			Code: branch.Code,
			Name: branch.Name,
		},
		Merchants: menuMerchants,
	}, nil
}

func (s *MerchantMenuService) MenuMerchant(merchantId uint) (*dto.MenusByMerchantResponse, error) {
	merchantMenus, err := s.merchantMenuRepo.FindByMerchantID(merchantId)
	if err != nil {
		return nil, err
	}

	catMap := make(map[uint]*dto.CategoryMenuResponse)

	for _, mm := range merchantMenus {
		cat := mm.Category
		menu := mm.Menu

		if _, ok := catMap[cat.ID]; !ok {
			catMap[cat.ID] = &dto.CategoryMenuResponse{
				ID:        cat.ID,
				Name:      cat.Name,
				SortOrder: cat.SortOrder,
				Items:     make([]dto.MenuWithPriceResponse, 0),
			}
		}

		catMap[cat.ID].Items = append(catMap[cat.ID].Items, dto.MenuWithPriceResponse{
			ID:          menu.ID,
			Name:        menu.Name,
			ImageURL:    menu.ImageURL,
			Description: menu.Description,
			Price:       mm.Price,
			Discount:    mm.Discount,
			SortOrder:   mm.SortOrder,
			FinalPrice:  mm.Price - mm.Discount,
		})
	}

	merchant, err := s.merchantRepo.FindByID(merchantId)
	if err != nil {
		return nil, err
	}

	resp := &dto.MenusByMerchantResponse{
		Merchant: dto.MerchantResponse{
			ID: merchant.ID,
			Branch: &dto.BranchResponse{
				ID:   merchant.Branch.ID,
				Code: merchant.Branch.Code,
				Name: merchant.Branch.Name,
			},
			Brand: &dto.BrandResponse{
				ID:   merchant.Brand.ID,
				Code: merchant.Brand.Code,
				Name: merchant.Brand.Name,
			},
			Channel: &dto.ChannelResponse{
				ID:   merchant.Channel.ID,
				Code: merchant.Channel.Code,
				Name: merchant.Channel.Name,
			},
		},
		Menu: dto.MenuMerchantResponse{
			Categories: make([]dto.CategoryMenuResponse, 0, len(catMap)),
		},
	}

	for _, cat := range catMap {
		sort.SliceStable(cat.Items, func(i, j int) bool {
			if cat.Items[i].SortOrder == cat.Items[j].SortOrder {
				return cat.Items[i].ID < cat.Items[j].ID
			}
			return cat.Items[i].SortOrder < cat.Items[j].SortOrder
		})
		resp.Menu.Categories = append(resp.Menu.Categories, *cat)
	}

	sort.SliceStable(resp.Menu.Categories, func(i, j int) bool {
		if resp.Menu.Categories[i].SortOrder == resp.Menu.Categories[j].SortOrder {
			return resp.Menu.Categories[i].ID < resp.Menu.Categories[j].ID
		}
		return resp.Menu.Categories[i].SortOrder < resp.Menu.Categories[j].SortOrder
	})

	return resp, nil
}
