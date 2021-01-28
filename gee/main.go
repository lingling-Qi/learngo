package main

import (
	"./geeH"
	"net/http"
)

func main() {
	r := geeH.New()

	r.GET("/", func(c *geeH.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *geeH.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *geeH.Context) {
		c.JSON(http.StatusOK, geeH.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
