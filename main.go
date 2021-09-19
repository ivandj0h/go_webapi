package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ivandi1980/controllers" // new
	"github.com/ivandi1980/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
