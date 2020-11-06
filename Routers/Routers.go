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
		v1.GET("/users/:id", Controllers.UserShow)
		v1.POST("/users", Controllers.UserStore)
		v1.PATCH("/users/:id", Controllers.UserUpdate)
		v1.DELETE("/users/:id", Controllers.UserDestroy)

		v1.GET("/roles", Controllers.RoleIndex)
		v1.POST("/roles", Controllers.RoleStore)
	}

	return r
}
