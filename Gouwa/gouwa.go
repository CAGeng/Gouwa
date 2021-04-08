package Gouwa

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine{
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc){
	log.Printf("Route %s - %s", method, pattern)
	engine.router.addRoute(method,pattern,handler)
}

//add GET request
func (engine *Engine) GET(pattern string,handler HandlerFunc){
	engine.addRoute("GET",pattern,handler)
}

//add POST request
func (engine * Engine) POST(pattern string,handler HandlerFunc){
	engine.addRoute("POST",pattern,handler)
}

//start HTTP server
func (engine * Engine) Run(addr string)(err error){
	return http.ListenAndServe(addr,engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	c := newContext(w,req)
	engine.router.handle(c)
}