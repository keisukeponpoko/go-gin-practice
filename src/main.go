package main

import (
	"main/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    controllers.Setup(router)

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
