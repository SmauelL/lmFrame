package main

import (
	"lm"
	"net/http"
)

func main() {
	r := lm.New()

	r.GET("/index", func(c *lm.Context) {
		c.HTML(http.StatusOK, "<h1>index</h1>")
	})

	v1 := r.Group("v1")
	{
		v1.GET("/", func(c *lm.Context) {
			c.HTML(http.StatusOK, "<h1>Hello lm</h1>")
		})

		v1.GET("/hello", func(c *lm.Context) {
			// expect /hello?name=lm
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("v2")
	{
		v2.GET("/hello/:name", func(c *lm.Context) {
			//expect /hello/lm
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *lm.Context) {
			c.JSON(http.StatusOK, lm.S{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
