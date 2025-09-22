package company

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	companyService *CompanyService
}

func NewCompanyHandler(factory *factory.Factory) *CompanyHandler {
	return &CompanyHandler{
		companyService: NewCompanyService(factory),
	}
}

func (h *CompanyHandler) Create(c *gin.Context) {
	var request dto.CompanyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.companyService.Create(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *CompanyHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request dto.CompanyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.companyService.Update(uint(id), &request); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, nil)
}

func (h *CompanyHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.companyService.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.NoContent(c)
}

func (h *CompanyHandler) Index(c *gin.Context) {
	companies, err := h.companyService.Index()
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, companies)
}

func (h *CompanyHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Find the company by ID
	company, err := h.companyService.Show(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, company)
}
