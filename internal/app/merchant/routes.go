package merchant

import (
	"sagala/factory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	merchantHandler := NewMerchantHandler(factory)

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/merchant", merchantHandler.Index)
			admin.GET("/merchant/:id", merchantHandler.Show)
			admin.POST("/merchant", merchantHandler.Create)
			admin.PUT("/merchant/:id", merchantHandler.Update)
			admin.DELETE("/merchant/:id", merchantHandler.Delete)
		}
	}
}
