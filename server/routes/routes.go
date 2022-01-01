package routes

import (
	"github.com/LOO2/business-remote-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		revenues := main.Group("revenue")
		{
			revenues.GET("/", controllers.ShowAllRevenues)
			revenues.GET("/:id", controllers.ShowRevenue)
			revenues.POST("/", controllers.CreateRevenue)
			revenues.PUT("/", controllers.UpdateRevenue)
			revenues.DELETE("/:id", controllers.DeleteRevenue)
		}

		cost_category := main.Group("cost_category")
		{
			cost_category.GET("/", controllers.ShowAllCostCategories)
			cost_category.GET("/:id", controllers.ShowCostCategory)
			cost_category.POST("/", controllers.CreateCostCategory)
			cost_category.PUT("/", controllers.UpdateCostCategory)
			cost_category.DELETE("/:id", controllers.DeleteCostCategory)
		}

		provider := main.Group("provider")
		{
			provider.GET("/", controllers.ShowAllProvider)
			provider.GET("/:id", controllers.ShowProviders)
			provider.POST("/", controllers.CreateProvider)
			provider.PUT("/", controllers.UpdateProvider)
			provider.DELETE("/:id", controllers.DeleteProvider)
		}
		cost := main.Group("cost")
		{
			cost.GET("/", controllers.ShowAllCost)
			cost.GET("/:id", controllers.ShowCost)
			cost.POST("/", controllers.CreateCost)
			cost.PUT("/", controllers.UpdateCost)
			cost.DELETE("/:id", controllers.DeleteCost)
		}
	}

	return router
}
