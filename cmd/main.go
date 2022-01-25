package main

import (
	"fmt"
	_ "log"

	"github.com/api-salto-theme/album"
)

// album represents data about a record album.
// type albumData struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albumDatas = []albumData{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	db := SqlConnect()
	defer db.Close()
	fmt.Println("Hello golang from docker!")
	album.Api(db)

	// albums slice to seed record album data.

	// router := gin.Default()
	// router.GET("/albums", getAlbumDatas)
	// router.Run("0.0.0.0:8080")
}

// getAlbums responds with the list of all albums as JSON.
// func getAlbumDatas(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albumDatas)
// }
