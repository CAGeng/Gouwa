package main

import (
	"./Gouwa"
	"log"
	"net/http"
	"time"
)


func onlyForV2() Gouwa.HandlerFunc {
	return func(c *Gouwa.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func onlyForV3() Gouwa.HandlerFunc {
	return func(c *Gouwa.Context) {
		// Start timer
		t := time.Now()
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := Gouwa.New()
	r.Use(Gouwa.Logger()) // global midlleware
	r.GET("/", func(c *Gouwa.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gouwa</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *Gouwa.Context) {
			// expect /hello/Gouwaktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

	}

	v3 := r.Group("/v3")
	v3.Use(onlyForV3())
	{
		v3.GET("/hello/:name", func(c *Gouwa.Context) {
			// expect /hello/Gouwaktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
