package brand

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	brandService *BrandService
}

func NewBrandHandler(factory *factory.Factory) *BrandHandler {
	return &BrandHandler{
		brandService: NewBrandService(factory),
	}
}

func (h *BrandHandler) Create(c *gin.Context) {
	var request dto.BrandRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.brandService.Create(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *BrandHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request dto.BrandRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.brandService.Update(uint(id), &request); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, nil)
}

func (h *BrandHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.brandService.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.NoContent(c)
}

func (h *BrandHandler) Index(c *gin.Context) {
	brands, err := h.brandService.Index()
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, brands)
}

func (h *BrandHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Find the brand by ID
	brand, err := h.brandService.Show(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, brand)
}
