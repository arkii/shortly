package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
var db *gorm.DB

func Init() {
	dbc, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer dbc.Close()
	db = dbc
}
*/

type Url struct {
	gorm.Model
	Key  string `gorm:"type:varchar(8);unique;not null"`
	Link string `gorm:"type:varchar(80000);not null"`
}

func Get(k string) (string, error) {
	fmt.Println(k)
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var url Url
	db.First(&url, "key = ?", k)
	fmt.Println(url)
	return url.Link, nil
}

func Shortly() {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Url{})

	// 创建
	db.Create(&Url{Key: "AAA", Link: "https://baidu.com/"})

	// 读取
	var url Url
	// db.First(&url, 1)                   // 查询id为1的product
	db.First(&url, "Key = ?", "AAA") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&url).Update("AAA", "https://google.com/")

	// 删除 - 删除product
	db.Delete(&url)
}
