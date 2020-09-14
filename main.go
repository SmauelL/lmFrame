package main

import (
	"lm"
	"net/http"
)

func main() {
	r := lm.Default()
	r.GET("/", func(c *lm.Context) {
		c.String(http.StatusOK, "Hello lm\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *lm.Context) {
		names := []string{"lm"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
