package main

import (
	"fmt"
	_ "log"

	"github.com/api-salto-theme/album"
)

func main() {
	db := SqlConnect()
	defer db.Close()
	fmt.Println("Hello golang from docker!")
	album.Api(db)
}
