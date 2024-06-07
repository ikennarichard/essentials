package main

import (
	"github.com/ikennarichard/essentials/db"
	"github.com/ikennarichard/essentials/lib"
	"github.com/ikennarichard/essentials/models"
)

func init() {
	lib.GetEnvVariables()
	db.ConnectToDb()
}

func main() {
	lib.DB.AutoMigrate(&models.Product{})
}