package day6

import (
	"net/http"
	"strings"
)

// HandlerFunc 便于处理请求
type HandlerFunc func(ctx *Context)

// Engine 是ownGin的核心
type Engine struct {
	*RouterGroup
	// 存储所有分组路由
	groups []*RouterGroup
	// 路由映射表
	router *router
}

// New 用来声明一个构造函数
func New() *Engine {
	// new与&的区别:
	// new不能进行初始化，默认零值；&可以进行初始化

	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(recovery())
	return engine
}

// Run 启动器，对http.ListenAndServe进行封装
func (e *Engine) Run(add ...string) error {
	if len(add) == 0 {
		return http.ListenAndServe(":80", e)
	}
	return http.ListenAndServe(add[0], e)

}

// 实现ServeHTTP接口，成为Handler
func (e *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 接收到一个请求时，先判断该请求需要经过哪些中间件
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(request.URL.Path, group.prefix) {
			// 然后将路由组中的中间件取出来添加到Context中
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	ctx := newContext(response, request)
	ctx.handlers = middlewares
	e.router.handle(ctx)
}
