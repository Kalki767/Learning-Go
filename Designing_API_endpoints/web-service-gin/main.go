package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//main function to route the requests and run the server
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)
	router.GET("/albums/:id", getAlbum)
    router.Run("localhost:8080")
}

//getAlbums function to get all the albums
func getAlbums(c *gin.Context){
    c.IndentedJSON(http.StatusOK, albums) //return the albums in json format
}

//addAlbum function to add a new album
func addAlbum(c *gin.Context){
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum. and return if there is any error
	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum) //return the new album in json format
}

//getAlbum function to get a specific album
func getAlbum(c *gin.Context){
	id := c.Param("id") //get the id from the request

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, album := range albums{
		if album.ID == id{
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	// If no album is found, return a 404 status with a JSON response.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
