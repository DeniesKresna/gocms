package Routers

import (
	"github.com/DeniesKresna/gocms/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/users", Controllers.UserIndex)
		v1.POST("/users", Controllers.UserStore)
		/*
		v1.POST("user", Controllers.AddNewUser)
		v1.GET("user/:id", Controllers.GetOneUser)
		v1.PUT("user/:id", Controllers.PutOneUser)
		v1.DELETE("user/:id", Controllers.DeleteUser)
		*/
	}

	return r
}
