package controllers

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
)

type UserController struct {
	Router *gin.Engine
	DbClient *sql.DB
}

func (controller *UserController) insertUser(c *gin.Context) {
	user := models.User{
		User: "name",
	}
	errs2 := user.Insert(context.Background(), controller.DbClient, boil.Infer())
	if errs2 != nil {
			panic(errs2.Error())
	}

	c.JSON(200, gin.H{
		"message": "users",
	})
}

func (controller *UserController) Setup() {
	v1 := controller.Router.Group("/user")
	{
		v1.GET("", controller.insertUser)
	}
}
