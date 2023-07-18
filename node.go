package goeasy

import (
	"github.com/daobin/goeasy/internal"
	"strings"
)

type handlerFunc func(c *Context)

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
	handlers handlerChain
	nType    internal.NodeType
	priority uint32
	children map[string]*node
}

// addRouteNode 添加路由节点（注册路由）
func (n *node) addRouteNode(fullPath string, handlers []handlerFunc) {
	if fullPath == "/" {
		n.handlers = handlers
		return
	}

	segments := strings.Split(fullPath, "/")
	for _, segment := range segments {
		if segment == "" {
			continue
		}

		if _, ok := n.children[segment]; !ok {
			n.children[segment] = &node{
				fullPath: internal.JoinPath(n.fullPath, segment),
				path:     segment,
				nType:    internal.NodeTypeNormal,
				children: map[string]*node{},
			}
		}
		n = n.children[segment]
	}

	n.handlers = handlers
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
