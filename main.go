package main

import (
	"./Gouwa"
	"net/http"
)


func main() {
	r := Gouwa.New()
	r.GET("/", func(c *Gouwa.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gouwa</h1>")
	})

	r.GET("/hello", func(c *Gouwa.Context) {
		// expect /hello?name=Gouwaktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *Gouwa.Context) {
		// expect /hello/Gouwaktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *Gouwa.Context) {
		c.JSON(http.StatusOK, Gouwa.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
