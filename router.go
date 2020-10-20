package main

import (
	"html/template"
	"net/http"

	"github.com/4ristodemus/log/helpers"
	"github.com/4ristodemus/log/services"
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func CreateRouter() *gin.Engine {
	Router.Static("/static", "static")
	Router.SetFuncMap(template.FuncMap{
		"renderMarkdown":  helpers.RenderMarkdown,
		"formatTimestamp": helpers.FormatTimestamp,
	})
	Router.LoadHTMLGlob("templates/*")

	Router.GET("/", func(c *gin.Context) {
		err, entries := services.AllEntries()
		if err != nil {
			c.AbortWithError(
				http.StatusInternalServerError,
				err,
			)
		} else {
			c.HTML(200, "index.tmpl", gin.H{
				"title":   "aristodemus / home",
				"entries": entries,
			})
		}
	})

	Router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	return Router
}
