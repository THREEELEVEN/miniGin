package day6

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	// 原始对象
	Response http.ResponseWriter
	Request  *http.Request
	// 请求信息
	Path   string
	Method string
	// url的path部分的参数
	UrlParams map[string]string
	// 相应状态码
	StatusCode int
	// 存储中间件及Handler
	handlers []HandlerFunc
	// 用于记录当前执行到第几个中间件
	index int
}

// 构造函数
func newContext(response http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		Response: response,
		Request:  request,
		Path:     request.URL.Path,
		Method:   request.Method,
		index:    -1,
	}
}

func (ctx *Context) Next() {
	ctx.index++
	handlersSliceLen := len(ctx.handlers)
	for ; ctx.index < handlersSliceLen; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Response.Header().Set(key, value)
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Request.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	// 给StatusCode字段赋值
	ctx.StatusCode = code
	// 给响应添加状态码
	ctx.Response.WriteHeader(code)
}

// 用于返回响应（原始内容）
func (ctx *Context) String(code int, format string, values ...any) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Response.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, obj any) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Response)
	err := encoder.Encode(obj)
	if err != nil {
		// StatusInternalServerError 服务器内部错误
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) GetUrlParam(key string) string {
	return ctx.UrlParams[key]
}

func (ctx *Context) Fail(code int, message string) {
	ctx.String(code, message)
}
