package company

import (
	"sagala/factory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, factory *factory.Factory) {
	companyHandler := NewCompanyHandler(factory)

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/company", companyHandler.Index)
			admin.GET("/company/:id", companyHandler.Show)
			admin.POST("/company", companyHandler.Create)
			admin.PUT("/company/:id", companyHandler.Update)
			admin.DELETE("/company/:id", companyHandler.Delete)
		}
	}
}
