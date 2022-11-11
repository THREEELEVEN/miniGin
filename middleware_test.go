package day5

import (
	"log"
	"net/http"
	"testing"
)

func middle1() HandlerFunc {
	return func(ctx *Context) {
		log.Println("Middleware: 1")
	}
}

func middle2() HandlerFunc {
	return func(ctx *Context) {
		log.Println("Middleware: 2")
	}
}

func TestMiddleware(t *testing.T) {
	engine := New()
	g1 := engine.NewGroup("/g1")
	g1.Use(middle1())
	{
		g1.GET("/name", func(ctx *Context) {
			ctx.String(http.StatusOK, "GroupRouter 1")
		})
	}

	g2 := engine.NewGroup("/g2")
	g2.Use(middle1(), middle2())
	{
		g2.GET("/age", func(ctx *Context) {
			ctx.String(http.StatusOK, "GroupRouter 2")
		})
	}

	engine.Run(":80")
}
