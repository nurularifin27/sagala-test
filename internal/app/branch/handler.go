package branch

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BranchHandler struct {
	branchService *BranchService
}

func NewBranchHandler(factory *factory.Factory) *BranchHandler {
	return &BranchHandler{
		branchService: NewBranchService(factory),
	}
}

func (h *BranchHandler) Create(c *gin.Context) {
	var request dto.BranchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.branchService.Create(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *BranchHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request dto.BranchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.branchService.Update(uint(id), &request); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, nil)
}

func (h *BranchHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.branchService.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.NoContent(c)
}

func (h *BranchHandler) Index(c *gin.Context) {
	var filter dto.BranchFilterRequest

	companyIDStr := c.Query("company_id")
	if companyIDStr != "" {
		companyID, parseErr := strconv.ParseUint(companyIDStr, 10, 64)
		if parseErr != nil {
			c.Error(parseErr)
			return
		}
		filter.CompanyID = uint(companyID)
	}
	branches, err := h.branchService.Index(&filter)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, branches)
}

func (h *BranchHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Find the branch by ID
	branch, err := h.branchService.Show(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, branch)
}
