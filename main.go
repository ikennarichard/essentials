package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/essentials/controllers"
	"github.com/ikennarichard/essentials/db"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	db.ConnectToDb()
}

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/", controllers.GetAllProducts)
	router.POST("/", controllers.AddProduct)

	router.Run()
}