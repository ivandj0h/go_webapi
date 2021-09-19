package main

import (
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
	router.POST("/albums", postAlbums)

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

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	album = append(album, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
