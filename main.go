package main

import (
    "github.com/gin-gonic/gin"
    "shortly/api"
)

func main() {
    r := gin.Default()
    api.InitRoutes(r)
    r.Run(":3000")
}
