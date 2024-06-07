package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/essentials/lib"
)

func init() {
	lib.GetEnvVariables()
}

func main() {
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"localhost"})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}