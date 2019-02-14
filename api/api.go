package api

import (
	"fmt"
	"github.com/arkii/shortly/dao"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Use(favicon.New("./static/favicon.ico"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/", func(c *gin.Context) {
		dao.Shortly()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/:uri", func(c *gin.Context) {
		uri := c.Param("uri")
		link, err := dao.Get(uri)
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect(http.StatusMovedPermanently, link)
		/*
			c.JSON(200, gin.H{
				"Link": link,
			})
		*/
	})
}
