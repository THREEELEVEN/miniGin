package day4

import (
	"testing"
)

func TestRouterGroup_GET(t *testing.T) {
	engine := New()
	g1 := engine.NewGroup("/g1")
	{
		g1.GET("/name", nil)
		g1.GET("/age", nil)
		g1.POST("/p", nil)
		g1.POST("/x", nil)
	}

	g2 := engine.NewGroup("/g2")
	{
		g2.GET("/id", nil)
	}
}
