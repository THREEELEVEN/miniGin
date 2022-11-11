package main

import (
	"OwnGin/day6"
	"net/http"
)

func main() {
	engine := day6.Default()

	g1 := engine.NewGroup("/g1")
	{

		g1.GET("/hello", func(ctx *day6.Context) {
			ctx.String(http.StatusOK, "Hello %s", ctx.Query("name"))
		})
	}

	g2 := engine.NewGroup("/g2")
	{
		g2.GET("/", func(ctx *day6.Context) {
			ctx.String(http.StatusOK, "Hello World by /g2/")
		})
	}

	engine.Run()
}
