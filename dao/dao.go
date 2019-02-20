package dao

import (
	"fmt"
	"github.com/arkii/shortly/short"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbPath = "data.db"

/*
var db *gorm.DB

func Init() {
	dbc, err := gorm.Open("sqlite3", dbPath)
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
	// fmt.Println(k)
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var url Url
	var l string
	if r := db.First(&url, "key = ?", k); r.Error != nil {
		fmt.Println("Key Not Found")
		l = "/"
	} else {
		l = url.Link
	}
	return l, nil
}

func New(l string) (string, error) {
	// fmt.Println(l)
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	keys := short.New(l)

	url := Url{Key: keys[0], Link: l}
	db.Create(&url)

	/* For duplicate keys, max 4 keys per 1 link
	url := Url{Key: "", Link: l}
	for _, k := range keys {
		url.Key = k
		if dbObj := db.Create(&url); dbObj.Error != nil {
			fmt.Println(dbObj.Error.Error())
		} else {
			break
		}
	}
	*/

	// fmt.Println(url)
	return url.Key, nil
}

func DbInit() {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Url{})
}
