package server

import (
	"net/http"

	"planning-poker/server/controllers"
	"planning-poker/server/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middleware.AuthMiddleware())

	v1 := router.Group("v1")
	{
		helloGroup := v1.Group("hello")
		{
			helloGroup.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"data": "hello world"})
			})
		}
		roomGroup := v1.Group("room")
		{
			room := new(controllers.RoomController)
			roomGroup.POST("/", room.Create)
			// roomGroup.GET("/:id", room.One)
			// roomGroup.GET("/", room.All)
			// roomGroup.PATCH("/:id", room.Update)
			// roomGroup.DELETE("/:id", room.Delete)
		}
	}
	return router

}
