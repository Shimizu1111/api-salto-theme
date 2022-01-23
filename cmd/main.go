package main

import (
	"fmt"
	_ "log"

	_ "github.com/Shimizu1111/api-salto-theme/api"
)

func main() {
	fmt.Println("Hello golang from docker!")
	api.getapi()
}
