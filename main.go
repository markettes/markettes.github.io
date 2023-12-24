package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load HTML templates from the "templates" directory
	router.LoadHTMLGlob("templates/*.tmpl")

	// Load static assets from the "assets" directory
	router.Static("/assets", "assets")

	// Define a route for the home page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "Identity Details",
			"name":    "Marcos Fernandez",
			"email":   "marferp2@inf.upv.es",
			"address": "Valencia (Spain🇪🇸)",
			"github":  "github.com/markettes",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
