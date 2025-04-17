package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/kataras/iris/v12"
	"net/http"
)

type Mux interface {
	HandleFunc(pattern string, handlerFn http.HandlerFunc)
	Get(pattern string, handlerFn http.HandlerFunc)
	Put(pattern string, handlerFn http.HandlerFunc)
	Delete(pattern string, handlerFn http.HandlerFunc)
	Routes() []chi.Route
	Handle(pattern string, handler http.Handler)
}

type irisMux struct {
	app *iris.Application
}

func NewIrisMux(app *iris.Application) Mux {
	return &irisMux{app: app}
}

func (i *irisMux) Handle(pattern string, handler http.Handler) {
	i.app.HandleMany("ALL", pattern, func(ctx iris.Context) {
		handler.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	})
}

func (i *irisMux) HandleFunc(pattern string, handlerFn http.HandlerFunc) {
	i.app.HandleMany("ALL", pattern, func(ctx iris.Context) {
		handlerFn(ctx.ResponseWriter(), ctx.Request())
	})
}

func (i *irisMux) Get(pattern string, handlerFn http.HandlerFunc) {
	i.app.Handle("GET", pattern, func(ctx iris.Context) {
		handlerFn(ctx.ResponseWriter(), ctx.Request())
	})
}

func (i *irisMux) Put(pattern string, handlerFn http.HandlerFunc) {
	i.app.Handle("PUT", pattern, func(ctx iris.Context) {
		handlerFn(ctx.ResponseWriter(), ctx.Request())
	})
}

func (i *irisMux) Delete(pattern string, handlerFn http.HandlerFunc) {
	i.app.Handle("DELETE", pattern, func(ctx iris.Context) {
		handlerFn(ctx.ResponseWriter(), ctx.Request())
	})
}
func (i *irisMux) Routes() []chi.Route {
	routes := i.app.GetRoutes()
	var list []chi.Route
	for _, route := range routes {
		list = append(list, chi.Route{
			SubRoutes: nil,
			Handlers:  nil,
			Pattern:   route.Path,
		})
	}
	return list
}
