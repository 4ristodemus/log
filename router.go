package main

import (
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func CreateRouter() *gin.Engine {
	Router.LoadHTMLGlob("templates/*")
	Router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "hello world",
		})
	})

	return Router
}
