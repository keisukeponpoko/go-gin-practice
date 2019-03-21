package controllers

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	user := UserController{Router: router}
	user.Setup()
}
