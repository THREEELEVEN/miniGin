package day6

/*
待优化：可以以组为单位创建不同的路由
*/

// RouterGroup 分组路由
type RouterGroup struct {
	// 前缀
	prefix string
	// 中间件
	middlewares []HandlerFunc
	// 整个框架的所有资源都是由Engine统一协调
	engine *Engine
}

// NewGroup 创建新路由分组
func (group *RouterGroup) NewGroup(prefix string) *RouterGroup {
	// 获取当前分组的Engine
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: prefix,
		engine: group.engine,
	}
	// 添加新分组
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 添加路由
func (group *RouterGroup) addRoute(requestMethod string, urlPart string, handler HandlerFunc) {
	//	// key由请求方法和静态路由地址构成：GET-/hello
	//	// 可以对同一请求地址根据请求方法做不同的处理
	pattern := group.prefix + urlPart
	// 打印注册的路由
	//log.Printf("Route %4s - %s", requestMethod, pattern)
	group.engine.router.addRoute(requestMethod, pattern, handler)
}

// GET 添加Get方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 添加Post方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Use 将Context中的中间件添加到路由组中
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
