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
	children map[string]*node
	isWild   bool
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

		isWild := segment[0] == ':' || segment[0] == '*'
		if len(segment) == 1 {
			// 路径片断不能只有一个通配符
			if isWild {
				return
			}
		} else {
			// 按字符循环查找匹配，通配符只能在片断中的第一位
			for _, char := range []byte(segment[1:]) {
				if char == ':' || char == '*' {
					return
				}
			}
		}

		child := n.matchChild(segment)
		if child == nil {
			child = &node{
				fullPath: internal.JoinPath(n.fullPath, segment),
				path:     segment,
				children: map[string]*node{},
				isWild:   isWild,
			}
			n.children[segment] = child
		}
		n = child

		// 只允许存在一个*通配符参数
		if segment[0] == '*' {
			break
		}
	}

	n.handlers = handlers
}

// matchChild 匹配子节点
func (n *node) matchChild(segment string) *node {
	if len(n.children) == 0 {
		return nil
	}

	for _, child := range n.children {
		if child.path == segment || child.isWild {
			return child
		}
	}

	return nil
}

// getChildNode 获取子节点（匹配路由）
func (n *node) getChildNode(segment string) *node {
	if len(n.children) == 0 {
		return nil
	}

	// 精确匹配优先
	child, ok := n.children[segment]
	if ok {
		return child
	}

	// 模糊匹配次之（路径参数匹配）
	for _, child = range n.children {
		if child.path[0] == ':' {
			return child
		}
	}
	for _, child = range n.children {
		if child.path[0] == '*' {
			return child
		}
	}

	return nil
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
