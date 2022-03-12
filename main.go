package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album struct data for record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.90},
	{ID: "2", Title: "African Giant", Artist: "Burna Boy", Price: 100.90},
	{ID: "3", Title: "This is Eminem", Artist: "Eminem", Price: 1560.90},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbum adds an album from JSON in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJson to bind the received JSON to newAlbum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add album
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	//loop over album to get album with the ID
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)

	router.Run("localhost:8080")
}
