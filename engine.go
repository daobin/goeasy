package goeasy

import (
	"github.com/daobin/goeasy/internal"
	"net/http"
	"sync"
)

// Engine 框架引擎
type Engine struct {
	router
	ctxPool sync.Pool
	trees   nodeTrees
}

// newContext 新建上下文
func (e *Engine) newContext() *context {
	return &context{}
}

// addRouteNode 添加路由节点
func (e *Engine) addRouteNode(httpMethod, absolutePath string, handlers []handlerFunc) {
	root := e.trees.get(httpMethod)
	if root == nil {
		root = &node{
			fullPath: "/",
			nType:    internal.NodeTypeNormal,
		}
		e.trees = append(e.trees, nodeTree{
			method: httpMethod,
			root:   root,
		})
	}

	root.addRoute(absolutePath, handlers)
}

// ServeHTTP 实现HTTP服务
func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

// 校验是否实现相关接口
var _ http.Handler = (*Engine)(nil)
