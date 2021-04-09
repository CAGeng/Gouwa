package main

import (
	"net/http"

	"./Gouwa"
)

func main() {
	r := Gouwa.Default()
	r.GET("/", func(c *Gouwa.Context) {
		c.String(http.StatusOK, "Hello Gouwaktutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *Gouwa.Context) {
		names := []string{"Gouwaktutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}