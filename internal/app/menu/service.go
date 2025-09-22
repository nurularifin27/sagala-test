package menu

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"

	"gorm.io/gorm"
)

type MenuService struct {
	menuRepo *repository.MenuRepository
}

func NewMenuService(factory *factory.Factory) *MenuService {
	return &MenuService{
		menuRepo: factory.MenuRepo,
	}
}

func (s *MenuService) Create(request *dto.MenuRequest) error {
	menu := models.Menu{
		Name:        request.Name,
		ImageURL:    request.ImageURL,
		Description: request.Description,
	}

	if err := s.menuRepo.Create(&menu); err != nil {
		return err
	}

	return nil
}

func (s *MenuService) Update(id uint, request *dto.MenuRequest) error {
	menu, err := s.menuRepo.FindByID(id)
	if err != nil {
		return err
	}

	menu.Name = request.Name
	menu.ImageURL = request.ImageURL
	menu.Description = request.Description

	return s.menuRepo.Update(&menu)
}

func (s *MenuService) Delete(id uint) error {
	return s.menuRepo.Delete(id)
}

func (s *MenuService) Index() ([]dto.MenuResponse, error) {
	var resp dto.MenuResponse
	menus, err := s.menuRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return resp.BuildListMenuResponse(menus), nil
}

func (s *MenuService) Show(id uint) (*dto.MenuResponse, error) {
	var resp dto.MenuResponse
	menu, err := s.menuRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return resp.BuildMenuResponse(menu), nil
}
