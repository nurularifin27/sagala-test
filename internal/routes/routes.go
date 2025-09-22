package routes

import (
	"sagala/factory"
	"sagala/internal/app/branch"
	"sagala/internal/app/brand"
	"sagala/internal/app/category"
	"sagala/internal/app/channel"
	"sagala/internal/app/company"
	"sagala/internal/app/menu"
	"sagala/internal/app/merchant"
	"sagala/internal/app/merchantmenu"
	"sagala/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	company.SetupRoutes(router, factory)
	branch.SetupRoutes(router, factory)
	brand.SetupRoutes(router, factory)
	channel.SetupRoutes(router, factory)
	category.SetupRoutes(router, factory)
	menu.SetupRoutes(router, factory)
	merchant.SetupRoutes(router, factory)
	merchantmenu.SetupRoutes(router, factory)

	router.NoRoute(func(c *gin.Context) {
		utils.NotFound(c, "Endpoint not found")
	})
}
