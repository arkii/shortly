package config

import (
    "encoding/json"
    "flag"
    "fmt"
    "os"
)

type Config struct {
    Database struct {
        Type string `json:"type"`
        Conn string `json:"conn"`
    } `json:"database"`
}

var Configuration Config

func init() {
    file := flag.String("c", "config.json", "Config file path")
    flag.Parse()
    Configuration = LoadConfiguration(file)
}

func LoadConfiguration(file *string) Config {
    var config Config
    configFile, err := os.Open(*file)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}
