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
	basePath    string
	handlers    []handlerFunc
	engine      *Engine
	isEngineNew bool
}

// 校验是否实现相关接口
var _ IRouter = (*router)(nil)

func (r *router) GET(relativePath string, handlers ...handlerFunc) IRouter {
	return r.handle(http.MethodGet, relativePath, handlers)
}

func (r *router) POST(relativePath string, handlers ...handlerFunc) IRouter {
	return r.handle(http.MethodPost, relativePath, handlers)
}

func (r *router) PUT(relativePath string, handlers ...handlerFunc) IRouter {
	return r.handle(http.MethodPut, relativePath, handlers)
}

func (r *router) DELETE(relativePath string, handlers ...handlerFunc) IRouter {
	return r.handle(http.MethodDelete, relativePath, handlers)
}

func (r *router) handle(httpMethod, relativePath string, handlers []handlerFunc) IRouter {
	absolutePath := r.calculateAbsolutePath(relativePath)
	handlers = r.mergeHandlers(handlers)
	r.engine.addRoute(httpMethod, absolutePath, handlers)

	return r.returnRouter()
}

func (r *router) calculateAbsolutePath(relativePath string) string {
	return internal.JoinPath(r.basePath, relativePath)
}

func (r *router) mergeHandlers(handlers []handlerFunc) []handlerFunc {
	finalSize := len(r.handlers) + len(handlers)
	// todo 可以通过finalSize控制最大数量的handlers

	mergeHandlers := make([]handlerFunc, finalSize)
	copy(mergeHandlers, r.handlers)
	copy(mergeHandlers[len(r.handlers):], handlers)

	return mergeHandlers
}

func (r *router) returnRouter() IRouter {
	if r.isEngineNew {
		return r.engine
	}

	return r
}

func (r *router) Group(relativePath string) *router {
	return &router{
		basePath:    r.calculateAbsolutePath(relativePath),
		handlers:    r.handlers,
		engine:      r.engine,
		isEngineNew: false,
	}
}
