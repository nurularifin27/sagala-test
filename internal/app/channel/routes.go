package channel

import (
	"sagala/factory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	channelHandler := NewChannelHandler(factory)

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/channel", channelHandler.Index)
			admin.GET("/channel/:id", channelHandler.Show)
			admin.POST("/channel", channelHandler.Create)
			admin.PUT("/channel/:id", channelHandler.Update)
			admin.DELETE("/channel/:id", channelHandler.Delete)
		}
	}
}
