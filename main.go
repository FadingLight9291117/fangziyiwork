package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Type string `json:"type" form:"type" binding:"required"`
	Ip   string `json:"ip" form:"ip" binding:"required"`
	Port int    `json:"port" form:"port" binding:"required"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "TCP/UDP 端口扫描程序",
		})
	})

	router.GET("/scan", func(c *gin.Context) {

		var data Data

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"msg": err})
		}

		c.JSON(200, gin.H{
			"msg": true,
		})
	})

	router.Run(":8080")

}
