package day3

import (
	"testing"
)

func TestRouter(t *testing.T) {
	router := newRouter()
	router.addRoute("GET", "/", nil)
	router.addRoute("GET", "/hello", nil)
	router.addRoute("GET", "/her", nil)
	router.addRoute("GET", "/:name", nil)
	router.addRoute("GET", "/:name/id", nil)
	router.addRoute("GET", "/:name/age", nil)
	router.addRoute("GET", "/:age", nil)

	router.getRoute("GET", "/her")
}

func TestSwitch(t *testing.T) {
}
