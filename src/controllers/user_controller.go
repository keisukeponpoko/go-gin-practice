package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Router *gin.Engine
}

func (controller *UserController) listUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users",
	})
}

func (controller *UserController) Setup() {
	v1 := controller.Router.Group("/user")
	{
		v1.GET("", controller.listUser)
	}
}
