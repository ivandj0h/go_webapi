package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// The 1st Route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"email":            "someone1@example.com",
			"password":         "mys3cr3T",
			"firstName":        "Firstname",
			"lastName":         "Lastname",
			"title":            "Mr",
			"primaryTelephone": "442071234567",
		})
	})

	router.Run()
}
