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

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
