package album

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Api(db *gorm.DB) {
	defer db.Close()
	fmt.Println("helloapi")
	fmt.Println("Hello golang from docker!")

	router := gin.Default()
	router.GET("/albums", getAlbumDatas)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", func(c *gin.Context) {
		postAlbumDatas(c, db)
	})
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

// postAlbums adds an album from JSON received in the request body.
func postAlbumDatas(c *gin.Context, db *gorm.DB) {
	//defer db.Close()

	var newAlbum albumData

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	dbInit(db)
	db.AutoMigrate(&albumData{})
	// Insert処理
	// TODO: ネットワークエラーかDB接続エラーかなどのエラーハンドリングを追加
	if err := db.Debug().Create(&albumData{ID: newAlbum.ID, Title: newAlbum.Title, Artist: newAlbum.Artist, Price: newAlbum.Price}).Error; err != nil {
		fmt.Println("insertのエラーです。", err)
	}
	// if err != nil {
	// 	fmt.Println("エラーです。", err)
	// }

	// Add the new album to the slice.
	albumDatas = append(albumDatas, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albumDatas {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// DBの初期化
func dbInit(db *gorm.DB) {
	db.Debug().AutoMigrate(&albumData{}) //構造体に基づいてテーブルを作成
	fmt.Println("テーブル挿入後")
}
