package goeasy

import "net/http"

type handlerFunc func(c *context)

type IRouter interface {
	GET(string, ...handlerFunc)
	POST(string, ...handlerFunc)
	PUT(string, ...handlerFunc)
	DELETE(string, ...handlerFunc)
}

type router struct {
	basePath string
	handlers []handlerFunc
	engine   *Engine
}

// 校验是否实现相关接口
var _ IRouter = (*router)(nil)

func (receiver *router) GET(relativePath string, handlers ...handlerFunc) {
	receiver.handler(http.MethodGet, relativePath, handlers)
}

func (receiver *router) POST(relativePath string, handlers ...handlerFunc) {
	receiver.handler(http.MethodPost, relativePath, handlers)
}

func (receiver *router) PUT(relativePath string, handlers ...handlerFunc) {
	receiver.handler(http.MethodPut, relativePath, handlers)
}

func (receiver *router) DELETE(relativePath string, handlers ...handlerFunc) {
	receiver.handler(http.MethodDelete, relativePath, handlers)
}

func (receiver *router) handler(httpMethod, relativePath string, handlers []handlerFunc) {

}

func (receiver *router) calculateAbsolutePath(relativePath string) {

}
