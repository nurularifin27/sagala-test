package category

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *CategoryService
}

func NewCategoryHandler(factory *factory.Factory) *CategoryHandler {
	return &CategoryHandler{
		categoryService: NewCategoryService(factory),
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var request dto.CategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.categoryService.Create(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request dto.CategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.categoryService.Update(uint(id), &request); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, nil)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.categoryService.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.NoContent(c)
}

func (h *CategoryHandler) Index(c *gin.Context) {
	categories, err := h.categoryService.Index()
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, categories)
}

func (h *CategoryHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Find the category by ID
	category, err := h.categoryService.Show(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, category)
}
