package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// The 1st Route (GET)
	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})
	router.GET("user/query", queryHandler)

	// The 2nd Route (POST)
	router.POST("/users", postUserHandler)

	// run the program
	router.Run()
}

// create rootHandler
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"email":            "juna@ivandjoh.online",
		"password":         "mys3cr3T",
		"firstName":        "Juna",
		"lastName":         "Djoh",
		"title":            "Mr",
		"primaryTelephone": "442071234567",
	})
}

// create helloHandler
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello World!",
	})
}

// Create queryHandler
func queryHandler(c *gin.Context) {
	nama := c.Query("nama")

	c.JSON(http.StatusOK, gin.H{
		"nama": nama,
	})
}

// create Struct
type UserInput struct {
	username string
	email    string
	age      int
	address  string
}

// Create postUserHandler
func postUserHandler(c *gin.Context) {
	var userInput UserInput

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"username": userInput.username,
		"email":    userInput.email,
		"age":      userInput.age,
		"address":  userInput.address,
	})
}
