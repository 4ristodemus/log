package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		fmt.Println("Navigated to root")

		c.HTML(200, "index.tmpl", gin.H{
			"title": "hellow world",
		})
	})

	router.Run()
}
