package brand

import (
	"sagala/factory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	brandHandler := NewBrandHandler(factory)

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/brand", brandHandler.Index)
			admin.GET("/brand/:id", brandHandler.Show)
			admin.POST("/brand", brandHandler.Create)
			admin.PUT("/brand/:id", brandHandler.Update)
			admin.DELETE("/brand/:id", brandHandler.Delete)
		}
	}
}
