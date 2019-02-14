package main

import (
	"github.com/arkii/shortly/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.InitRoutes(r)
	r.Run(":3000")
}
