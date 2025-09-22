package menu

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	menuService *MenuService
}

func NewMenuHandler(factory *factory.Factory) *MenuHandler {
	return &MenuHandler{
		menuService: NewMenuService(factory),
	}
}

func (h *MenuHandler) Create(c *gin.Context) {
	var request dto.MenuRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.menuService.Create(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *MenuHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request dto.MenuRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.menuService.Update(uint(id), &request); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, nil)
}

func (h *MenuHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.menuService.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.NoContent(c)
}

func (h *MenuHandler) Index(c *gin.Context) {
	menus, err := h.menuService.Index()
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, menus)
}

func (h *MenuHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Find the menu by ID
	menu, err := h.menuService.Show(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, menu)
}
