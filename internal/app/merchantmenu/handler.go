package merchantmenu

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MerchantMenuHandler struct {
	merchantMenuService *MerchantMenuService
}

func NewMerchantMenuHandler(factory *factory.Factory) *MerchantMenuHandler {
	return &MerchantMenuHandler{
		merchantMenuService: NewMerchantMenuService(factory),
	}
}

func (h *MerchantMenuHandler) Create(c *gin.Context) {
	var request dto.MerchantMenuRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.merchantMenuService.Create(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *MerchantMenuHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request dto.MerchantMenuRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.merchantMenuService.Update(uint(id), &request); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, nil)
}

func (h *MerchantMenuHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.merchantMenuService.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.NoContent(c)
}

func (h *MerchantMenuHandler) Index(c *gin.Context) {
	var filter dto.MerchantMenuFilterRequest

	merchantIDStr := c.Query("merchant_id")
	if merchantIDStr != "" {
		merchantID, parseErr := strconv.ParseUint(merchantIDStr, 10, 64)
		if parseErr != nil {
			c.Error(parseErr)
			return
		}
		filter.MerchantID = uint(merchantID)
	}

	merchantMenus, err := h.merchantMenuService.Index(&filter)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, merchantMenus)
}

func (h *MerchantMenuHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Find the merchantMenu by ID
	merchantMenu, err := h.merchantMenuService.Show(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, merchantMenu)
}

func (h *MerchantMenuHandler) UpdatePrice(c *gin.Context) {
	var request dto.MerchantMenuPriceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.merchantMenuService.UpdatePrice(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *MerchantMenuHandler) MenuBranch(c *gin.Context) {
	branchIdParam := c.Param("branchId")
	branchId, err := strconv.ParseUint(branchIdParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	menus, err := h.merchantMenuService.MenuBranch(uint(branchId))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, menus)
}

func (h *MerchantMenuHandler) MenuMerchant(c *gin.Context) {
	merchantIdParam := c.Param("merchantId")
	merchantId, err := strconv.ParseUint(merchantIdParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	menus, err := h.merchantMenuService.MenuMerchant(uint(merchantId))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, menus)
}
