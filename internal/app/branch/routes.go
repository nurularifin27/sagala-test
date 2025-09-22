package branch

import (
	"sagala/factory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	branchHandler := NewBranchHandler(factory)

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/branch", branchHandler.Index)
			admin.GET("/branch/:id", branchHandler.Show)
			admin.POST("/branch", branchHandler.Create)
			admin.PUT("/branch/:id", branchHandler.Update)
			admin.DELETE("/branch/:id", branchHandler.Delete)
		}
	}
}
