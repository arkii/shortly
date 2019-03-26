package dao

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "shortly/config"
    "shortly/short"
)

func DbConnect() *gorm.DB {
    db, err := gorm.Open(config.Configuration.Database.Type, config.Configuration.Database.Conn)
    if err != nil {
        panic("failed to connect database")
    }
    return db
}

type Url struct {
    gorm.Model
    Key  string `gorm:"type:varchar(8);unique;not null"`
    Link string `gorm:"type:varchar(80000);not null"`
}

func Get(k string) (string, error) {
    var url Url
    var l string
    db := DbConnect()
    defer db.Close()
    if r := db.First(&url, "key = ?", k); r.Error != nil {
        fmt.Println("Key Not Found")
        l = "/"
    } else {
        l = url.Link
    }
    return l, nil
}

func New(l string) (string, error) {
    keys := short.New(l)

    url := Url{Key: keys[0], Link: l}
    db := DbConnect()
    defer db.Close()
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

    return url.Key, nil
}

func DbInit() {
    db := DbConnect()
    defer db.Close()
    // Migrate the schema
    db.AutoMigrate(&Url{})
}
