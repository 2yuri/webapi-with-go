package routes

import (
	"github.com/LOO2/business-remote-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		revenues := main.Group("revenue")
		{
			revenues.GET("/", controllers.ShowAllRevenues)
			revenues.GET("/:id", controllers.ShowRevenue)
			revenues.POST("/", controllers.CreateRevenue)
			revenues.DELETE("/:id", controllers.DeleteRevenue)
		}
	}

	return router
}
