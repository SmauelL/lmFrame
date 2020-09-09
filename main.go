package main

import (
	"lm"
	"net/http"
)

func main() {
	r := lm.New()

	r.GET("/", func(c *lm.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Lm</h1>")
	})

	r.GET("/hello", func(c *lm.Context) {
		// expect /hello?name=lm
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *lm.Context) {
		c.JSON(http.StatusOK, lm.S{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
