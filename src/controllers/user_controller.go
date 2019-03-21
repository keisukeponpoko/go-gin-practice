package controllers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
	"main/models"
)

type UserController struct {
	Router     *gin.Engine
	Connection *sql.DB
}

func (controller *UserController) insertUser(c *gin.Context) {
	user := models.User{
		User: "name",
	}
	errs2 := user.Insert(context.Background(), controller.Connection, boil.Infer())
	if errs2 != nil {
		panic(errs2.Error())
	}

	users, err := models.Users().All(context.Background(), controller.Connection)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, gin.H{
		"message": users,
	})
}

func (controller *UserController) Setup() {
	v1 := controller.Router.Group("/user")
	{
		v1.GET("", controller.insertUser)
	}
}
