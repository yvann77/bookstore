package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/your-username/bookstore/handlers"
)

func SetupBookRoutes(router *gin.Engine) {
	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("/", handlers.GetBooks)
		bookRoutes.POST("/", handlers.PostBooks)
		bookRoutes.GET("/:id", handlers.GetBookByID)
	}
}
