package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, db *sql.DB) {
	user := UserController{Router: router, Connection: db}
	user.Setup()
}
