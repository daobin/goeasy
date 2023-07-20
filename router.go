package goeasy

import (
	"github.com/daobin/goeasy/internal"
	"net/http"
)

type IRouter interface {
	Use(middlewares ...handlerFunc) IRouter
	GET(relativePath string, handlers ...handlerFunc) IRouter
	POST(relativePath string, handlers ...handlerFunc) IRouter
	PUT(relativePath string, handlers ...handlerFunc) IRouter
	DELETE(relativePath string, handlers ...handlerFunc) IRouter
}

type router struct {
	basePath    string
	handlers    handlerChain
	engine      *Engine
	isEngineNew bool
}

// 校验是否实现相关接口
var _ IRouter = (*router)(nil)

func (r *router) Use(middlewares ...handlerFunc) IRouter {
	r.handlers = append(r.handlers, middlewares...)
	return r.returnRouter()
}

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

// handle 路由节点添加处理
func (r *router) handle(httpMethod, relativePath string, handlers handlerChain) IRouter {
	absolutePath := r.calculateAbsolutePath(relativePath)
	handlers = r.mergeHandlers(handlers)
	r.engine.addRouteNode(httpMethod, absolutePath, handlers)

	return r.returnRouter()
}

// calculateAbsolutePath 返回完整的请求路径
func (r *router) calculateAbsolutePath(relativePath string) string {
	return internal.JoinPath(r.basePath, relativePath)
}

// mergeHandlers 合并多个处理函数
func (r *router) mergeHandlers(handlers []handlerFunc) []handlerFunc {
	finalSize := len(r.handlers) + len(handlers)
	// todo 可以通过finalSize控制最大数量的handlers

	mergeHandlers := make([]handlerFunc, finalSize)
	copy(mergeHandlers, r.handlers)
	copy(mergeHandlers[len(r.handlers):], handlers)

	return mergeHandlers
}

// returnRouter 返回请求路由
func (r *router) returnRouter() IRouter {
	if r.isEngineNew {
		return r.engine
	}

	return r
}

// Group 返回新建路由组
func (r *router) Group(relativePath string) *router {
	return &router{
		basePath:    r.calculateAbsolutePath(relativePath),
		handlers:    r.handlers,
		engine:      r.engine,
		isEngineNew: false,
	}
}
