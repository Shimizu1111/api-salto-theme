package album

import (
	"fmt"
	_ "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getapi() {
	fmt.Println("helloapi")

	fmt.Println("Hello golang from docker!")
	router := gin.Default()
	router.GET("/albums", getAlbumDatas)
	router.Run("0.0.0.0:8080")
}

// album represents data about a record album.
type albumData struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albumDatas = []albumData{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbumDatas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumDatas)
}
