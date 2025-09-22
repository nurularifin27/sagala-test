package merchant

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	merchantService *MerchantService
}

func NewMerchantHandler(factory *factory.Factory) *MerchantHandler {
	return &MerchantHandler{
		merchantService: NewMerchantService(factory),
	}
}

func (h *MerchantHandler) Create(c *gin.Context) {
	var request dto.MerchantRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.merchantService.Create(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *MerchantHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request dto.MerchantRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.merchantService.Update(uint(id), &request); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, nil)
}

func (h *MerchantHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.merchantService.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.NoContent(c)
}

func (h *MerchantHandler) Index(c *gin.Context) {
	merchants, err := h.merchantService.Index()
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, merchants)
}

func (h *MerchantHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Find the merchant by ID
	merchant, err := h.merchantService.Show(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, merchant)
}
