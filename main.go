package main

import (
	"fmt"
	"html/template"
	"lm"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := lm.New()
	r.Use(lm.Logger())
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "lm", Age: 25}
	stu2 := &student{Name: "ljx", Age: 26}
	r.GET("/", func(c *lm.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *lm.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", lm.S{
			"title":  "lm",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *lm.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", lm.S{
			"title": "lm",
			"now":   time.Date(2020, 9, 11, 15, 20, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
