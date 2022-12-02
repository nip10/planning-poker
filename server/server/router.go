package server

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.Use(middlewares.AuthMiddleware())

	// r.GET("/", func(c *gin.Context) {
	//   c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	v1 := router.Group("v1")
	{
		// userGroup := v1.Group("user")
		// {
		// 	user := new(controllers.UserController)
		// 	userGroup.GET("/:id", user.Retrieve)
		// }
	}
	return router

}
