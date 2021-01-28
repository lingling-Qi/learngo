package geeH

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	route *router
}

func New() *Engine {
	return &Engine{newRoute()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.route.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.route.handle(c)
}
