package main

import (
	"lm"
	"log"
	"net/http"
	"time"
)

func onlyForV2() lm.HandlerFunc {
	return func(c *lm.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := lm.New()
	r.Use(lm.Logger()) // global middleware
	r.GET("/", func(c *lm.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Lm</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *lm.Context) {
			//expect /hello/lm
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
