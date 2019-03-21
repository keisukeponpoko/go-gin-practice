package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db, errs := sql.Open("mysql", "root:root@tcp(mysql:3306)/test")
	if errs != nil {
		panic(errs.Error())
	}
	defer db.Close()

	router := gin.Default()
	controllers.Setup(router, db)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	err := router.Run(":8083")
	if err != nil {
		panic(err)
	}
}
