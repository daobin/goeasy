package goeasy

import (
	"github.com/daobin/goeasy/internal"
)

type handlerFunc func(c *context)

type handlerChain []handlerFunc

func (h handlerChain) Last() handlerFunc {
	if length := len(h); length > 0 {
		return h[length]
	}

	return nil
}

type node struct {
	fullPath string
	path     string
	indices  string
	handlers handlerChain
	nType    internal.NodeType
	priority uint32
	children []*node
	isEnd    bool
}

// addRouteNode 添加路由节点（注册路由）
func (n *node) addRouteNode(fullPath string, handlers []handlerFunc) {

}

type nodeTree struct {
	method string
	root   *node
}

type nodeTrees []nodeTree

// get 返回指定请求方式的Root节点
func (nts nodeTrees) get(method string) *node {
	for _, tree := range nts {
		if tree.method == method {
			return tree.root
		}
	}

	return nil
}
