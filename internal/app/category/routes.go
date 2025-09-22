package category

import (
	"sagala/factory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	categoryHandler := NewCategoryHandler(factory)

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/category", categoryHandler.Index)
			admin.GET("/category/:id", categoryHandler.Show)
			admin.POST("/category", categoryHandler.Create)
			admin.PUT("/category/:id", categoryHandler.Update)
			admin.DELETE("/category/:id", categoryHandler.Delete)
		}
	}
}
