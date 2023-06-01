package goeasy

import "github.com/daobin/goeasy/internal"

type node struct {
	fullPath string
	handlers []handlerFunc
	nType    internal.NodeType
	priority uint32
	children []*node
}

// addRoute 添加路由节点
func (n *node) addRoute(path string, handlers []handlerFunc) {
	// 匹配新路径与当前路径的相同前缀最大长度
	longest := internal.CommonPrefixLongest(path, n.fullPath)

	// 新路径与当前路径一致
	if longest == len(path) && longest == len(n.fullPath) {
		if len(n.handlers) == 0 {
			n.handlers = handlers
		}

		return
	}

	if len(path) < len(n.fullPath) {
		child := node{
			fullPath: n.fullPath,
			handlers: n.handlers,
			priority: n.priority + 1,
			children: n.children,
		}

		n.children = []*node{&child}
		n.fullPath = path
		n.handlers = handlers

		return
	}

	//fullPath := path
	//
	//n.priority++
	//if n.path == "" && len(n.children) == 0 {
	//	n.insertWildChild(path, fullPath, handlers)
	//	n.nType = internal.NodeTypeNormal
	//	return
	//}
	//
	//parentFullPathIndex := 0
	//
	//for {
	//	// path 新传入的路径
	//	// n.path 当前节点的路径
	//	longest := internal.CommonPrefixLongest(path, n.path)
	//
	//	if longest < len(n.path) {
	//		child := node{
	//			path:     n.path[longest:],
	//			fullPath: n.fullPath,
	//			handlers: n.handlers,
	//			priority: n.priority - 1,
	//			children: n.children,
	//		}
	//
	//		n.children = []*node{&child}
	//		n.path = path[:longest]
	//		n.fullPath = fullPath[:parentFullPathIndex+longest]
	//		n.handlers = handlers
	//	}
	//
	//	if longest < len(path) {
	//		path = path[longest:]
	//		char := path[0]
	//
	//		if n.nType == internal.NodeTypeParam && char == '/' && len(n.children) == 1 {
	//			parentFullPathIndex += len(n.path)
	//
	//			n = n.children[0]
	//			n.priority++
	//			continue
	//		}
	//	}
	//}

}

func (n *node) insertWildChild(path, fullPath string, handlers []handlerFunc) {
	for {
		wildcard, idx, valid := internal.FindPathWildcard(path)
		if idx < 0 {
			break
		}

		if !valid {
			panic(internal.MergeString("路径通配符错误：", fullPath, " >> ", wildcard))
		}

		if wildcard[0] == ':' {
			if idx > 0 {
				//n.path = path[:idx]
				path = path[idx:]
			}

			child := node{
				//path:     wildcard,
				fullPath: fullPath,
				nType:    internal.NodeTypeParam,
			}

			n.children = append(n.children, &child)
			n = &child
			n.priority++

			if len(wildcard) < len(path) {
				path = path[len(wildcard):]
				child = node{fullPath: fullPath, priority: 1}
				n.children = append(n.children, &child)
				n = &child
				continue
			}

			n.handlers = handlers
			return
		}

	}

	//n.path = path
	//n.fullPath = fullPath
	//n.handlers = handlers
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
