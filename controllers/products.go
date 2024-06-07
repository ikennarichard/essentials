package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/essentials/lib"
	"github.com/ikennarichard/essentials/models"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product

	result := lib.DB.Find(&products)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Failed to retrieve products",
		})
		return
	}

	c.JSON(200, products)
}