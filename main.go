package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tangren1998/ddd-cheanArch/user/infra/api"
)

func main() {
	r := gin.Default()
	r.GET("/user/:id", api.GetUserById)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}