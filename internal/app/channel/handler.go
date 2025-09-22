package channel

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChannelHandler struct {
	channelService *ChannelService
}

func NewChannelHandler(factory *factory.Factory) *ChannelHandler {
	return &ChannelHandler{
		channelService: NewChannelService(factory),
	}
}

func (h *ChannelHandler) Create(c *gin.Context) {
	var request dto.ChannelRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.channelService.Create(&request); err != nil {
		c.Error(err)
		return
	}

	utils.Created(c, nil)
}

func (h *ChannelHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request dto.ChannelRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.channelService.Update(uint(id), &request); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, nil)
}

func (h *ChannelHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.channelService.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.NoContent(c)
}

func (h *ChannelHandler) Index(c *gin.Context) {
	channels, err := h.channelService.Index()
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, channels)
}

func (h *ChannelHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Find the channel by ID
	channel, err := h.channelService.Show(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, channel)
}
