package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		timeoff, ok := TimesOff[id]
		if !ok {
			c.String(http.StatusNotFound, "404-Page not Found")
			return
		}
		c.HTML(http.StatusOK, "vacation-overview.html",
			map[string]interface{}{
				"TimesOff": timeoff,
			})

	})

	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})
	r.Static("/public", "./public")
	return r
}
