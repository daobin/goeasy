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

func (receiver *Engine) allocateContext() *context {
	return &context{}
}

func (receiver *Engine) addRoute(httpMethod, absolutePath string, handlers []handlerFunc) {
	root := receiver.trees.get(httpMethod)
	if root == nil {
		root = &node{
			fullPath: "/",
			nType:    internal.NodeTypeRoot,
		}
		receiver.trees = append(receiver.trees, nodeTree{
			method: httpMethod,
			root:   root,
		})
	}

	root.addRoute(absolutePath, handlers)
}

func (receiver *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

// 校验是否实现相关接口
var _ http.Handler = (*Engine)(nil)
