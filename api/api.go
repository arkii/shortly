package api

import (
	"fmt"
	"github.com/arkii/shortly/dao"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

type Form struct {
	Link string `form:"link" json:"link" xml:"link"  binding:"required"`
}

func InitRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Use(favicon.New("./static/favicon.ico"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/", func(c *gin.Context) {
		var form Form
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		key, err := dao.New(form.Link)
		if err != nil {
			fmt.Println(err)
		}
		/*
			c.JSON(200, gin.H{
				"Key":  key,
				"Link": form.Link,
			})
		*/
		c.HTML(http.StatusOK, "new.html", gin.H{
			"Key":  key,
			"Link": form.Link,
		})
	})
	r.GET("/:uri", func(c *gin.Context) {
		uri := c.Param("uri")
		link, err := dao.Get(uri)
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect(http.StatusMovedPermanently, link)
	})
	/* For init db
	r.PUT("/", func(c *gin.Context) {
		dao.DbInit()
		c.JSON(200, gin.H{
			"message": "done",
		})
	})
	*/
}
