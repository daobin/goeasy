package goeasy

import (
	"net/http"
	"sync"
)

// Engine 框架引擎
type Engine struct {
	router
	ctxPool sync.Pool
	trees   nodeTrees
}

func (e *Engine) allocateContext() *context {
	return &context{}
}

func (e *Engine) addRoute(httpMethod, absolutePath string, handlers []handlerFunc) {
	root := e.trees.get(httpMethod)
	if root == nil {
		root = &node{
			fullPath: "/",
		}
		e.trees = append(e.trees, nodeTree{
			method: httpMethod,
			root:   root,
		})
	}

	root.addRoute(absolutePath, handlers)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

// 校验是否实现相关接口
var _ http.Handler = (*Engine)(nil)
