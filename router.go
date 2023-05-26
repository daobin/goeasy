package goeasy

import (
	"github.com/daobin/goeasy/internal"
	"net/http"
)

type handlerFunc func(c *context)

type IRouter interface {
	GET(string, ...handlerFunc) IRouter
	POST(string, ...handlerFunc) IRouter
	PUT(string, ...handlerFunc) IRouter
	DELETE(string, ...handlerFunc) IRouter
}

type router struct {
	basePath string
	handlers []handlerFunc
	engine   *Engine
}

// 校验是否实现相关接口
var _ IRouter = (*router)(nil)

func (receiver *router) GET(relativePath string, handlers ...handlerFunc) IRouter {
	return receiver.handler(http.MethodGet, relativePath, handlers)
}

func (receiver *router) POST(relativePath string, handlers ...handlerFunc) IRouter {
	return receiver.handler(http.MethodPost, relativePath, handlers)
}

func (receiver *router) PUT(relativePath string, handlers ...handlerFunc) IRouter {
	return receiver.handler(http.MethodPut, relativePath, handlers)
}

func (receiver *router) DELETE(relativePath string, handlers ...handlerFunc) IRouter {
	return receiver.handler(http.MethodDelete, relativePath, handlers)
}

func (receiver *router) handler(httpMethod, relativePath string, handlers []handlerFunc) IRouter {
	absolutePath := receiver.calculateAbsolutePath(relativePath)
	handlers = receiver.mergeHandlers(handlers)
	receiver.engine.addRoute(httpMethod, absolutePath, handlers)

	return receiver
}

func (receiver *router) calculateAbsolutePath(relativePath string) string {
	return internal.JoinPath(receiver.basePath, relativePath)
}

func (receiver *router) mergeHandlers(handlers []handlerFunc) []handlerFunc {
	finalSize := len(receiver.handlers) + len(handlers)
	// todo 可以通过finalSize控制最大数量的handlers

	mergeHandlers := make([]handlerFunc, finalSize)
	copy(mergeHandlers, receiver.handlers)
	copy(mergeHandlers[len(receiver.handlers):], handlers)

	return mergeHandlers
}

func (receiver *router) Group(relativePath string) *router {
	return &router{
		basePath: receiver.calculateAbsolutePath(relativePath),
		handlers: receiver.handlers,
		engine:   receiver.engine,
	}
}
