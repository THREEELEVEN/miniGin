package day6

import (
	"net/http"
	"strings"
)

type router struct {
	/*
			roots 存储每种请求方式(GET,POST)的前缀树根节点
					GET				   	POST
					/ \					/  \
		           g1 g2			   p1  p2

	*/
	roots map[string]*node

	/*
		k: 由请求路径和请求方法构成(RequestMethod-URL)
		v: 相应的处理方法(handler)
		GET-/g/name -> func
		POST-/p/id -> func
	*/
	handlers map[string]HandlerFunc
}

// 构造函数
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 解析URL的path参数部分，以“/”切分为slice
func parseUrlPath(urlPath string) []string {
	urlParts := strings.Split(urlPath, "/")
	resultParts := make([]string, 0)

	for _, part := range urlParts {
		// 将不为空的part添加到结果集中
		if part != "" {
			resultParts = append(resultParts, part)
			// 如果出现通配符‘*’则直接退出循环
			if part[0] == '*' {
				break
			}
		}
	}
	return resultParts
}

func (r *router) addRoute(requestMethod string, urlPath string, handler HandlerFunc) {
	parts := parseUrlPath(urlPath)
	// Get-/hello
	key := requestMethod + "-" + urlPath

	// 查看roots中是否存在相应的RequestMethod(GET, POST)
	_, ok := r.roots[requestMethod]
	// 如果roots中没有相应的RequestMethod
	if !ok {
		// 则创建一个新的节点
		r.roots[requestMethod] = &node{}
	}
	// 然后将该路由part插入到这个新节点下
	r.roots[requestMethod].insert(urlPath, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(requestMethod string, urlPath string) (*node, map[string]string) {
	willBeSearchedParts := parseUrlPath(urlPath)

	params := make(map[string]string)

	// 在路由树中查找是否有对应的路由
	root, ok := r.roots[requestMethod]
	if !ok {
		return nil, nil
	}

	childNode := root.search(willBeSearchedParts, 0)

	// 如果在路由树中找到了对应的路由
	if childNode != nil {
		parts := parseUrlPath(childNode.urlPath)

		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = willBeSearchedParts[index]
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(willBeSearchedParts[index:], "/")
				break
			}
		}
		return childNode, params
	}
	return nil, nil
}

func (r *router) handle(ctx *Context) {
	childNode, params := r.getRoute(ctx.Method, ctx.Path)

	if childNode != nil {
		ctx.UrlParams = params
		key := ctx.Method + "-" + childNode.urlPath
		ctx.handlers = append(ctx.handlers, r.handlers[key])
	} else {
		ctx.handlers = append(ctx.handlers, func(ctx *Context) {
			ctx.String(http.StatusNotFound, "404 Not Found: %s\n", ctx.Path)
		})
	}
	ctx.Next()
}
