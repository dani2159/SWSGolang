package router

import (
	bookControllers "SWSGolang/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/books", bookControllers.GetAllBooks)
		v1.POST("/books", bookControllers.CreateBook)
		v1.GET("/books/:bookID", bookControllers.GetBook)
		v1.PUT("/books/:bookID", bookControllers.UpdateBook)
		v1.DELETE("/books/:bookID", bookControllers.DeleteBook)
	}

	router.Run(":8080")
	return router
}