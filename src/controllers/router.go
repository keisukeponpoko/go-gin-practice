package controllers

import (
	"github.com/gin-gonic/gin"
	"database/sql"
)

func Setup(router *gin.Engine, db *sql.DB) {
	user := UserController{Router: router, DbClient: db}
	user.Setup()
}
