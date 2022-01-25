package main

import (
	// 追加でインポート
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	// _ でインポートすることで使用しなくてもコンパイルエラーにならない。
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBへの接続
func SqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	// 作成例: go_test:password@tcp(db:3306)/go_database?charset=utf8&parseTime=true&loc=Asia%2FTokyo

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return db
}
