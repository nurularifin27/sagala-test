package merchantmenu

import (
	"sagala/factory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	merchantMenuHandler := NewMerchantMenuHandler(factory)

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/merchant-menu", merchantMenuHandler.Index)
			admin.GET("/merchant-menu/:id", merchantMenuHandler.Show)
			admin.POST("/merchant-menu", merchantMenuHandler.Create)
			admin.PUT("/merchant-menu/:id", merchantMenuHandler.Update)
			admin.DELETE("/merchant-menu/:id", merchantMenuHandler.Delete)

			admin.PUT("/merchant-menu/price", merchantMenuHandler.UpdatePrice)
		}

		v1.GET("/merchant-menu/branch/:branchId", merchantMenuHandler.MenuBranch)
		v1.GET("/merchant-menu/merchant/:merchantId", merchantMenuHandler.MenuMerchant)
	}
}
