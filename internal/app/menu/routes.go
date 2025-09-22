package menu

import (
	"sagala/factory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	menuHandler := NewMenuHandler(factory)

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/menu", menuHandler.Index)
			admin.GET("/menu/:id", menuHandler.Show)
			admin.POST("/menu", menuHandler.Create)
			admin.PUT("/menu/:id", menuHandler.Update)
			admin.DELETE("/menu/:id", menuHandler.Delete)
		}
	}
}
