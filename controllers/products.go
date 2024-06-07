package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/essentials/lib"
	"github.com/ikennarichard/essentials/models"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product

	result := lib.DB.Find(&products)

	if result.Error != nil {
		log.Fatal("Couldnt get products")
		c.JSON(400, gin.H{
			"error": "Failed to retrieve products",
		})
		return
	}

	c.JSON(200, products)
}

func AddProduct(c *gin.Context) {
	var product models.Product

	var body struct { //structure of the request body
		Name string `json:"name"`
		Description string `json:"description"`
		Price float64 `json:"price"`
		ImageURL string `json:"image_url"`
	}

	c.Bind(&body)

	product = models.Product{Name:body.Name, Description:body.Description, Price: body.Price, ImageURL: body.ImageURL,}

	result := lib.DB.Create(&product)

	if result.Error != nil {
		log.Fatal("Product could not be created")
		c.JSON(500, gin.H{
			"error": "Failed to retrieve products",
		})
		return
	}

	c.JSON(201, product)
}

