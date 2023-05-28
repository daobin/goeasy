package goeasy

import "github.com/daobin/goeasy/internal"

type node struct {
	path     string
	fullPath string
	handlers []handlerFunc
	nType    internal.NodeType
	priority uint32
	children []*node
}

func (n *node) addRoute(path string, handlers []handlerFunc) {
	fullPath := path

	n.priority++
	if n.path == "" && len(n.children) == 0 {
		n.insertWildChild(path, fullPath, handlers)
		n.nType = internal.NodeTypeRoot
		return
	}

	parentFullPahtIndex := 0

	for {
		longest := internal.CommonPrefixLongest(path, n.path)

		if longest < len(n.path) {
			child := node{
				path:     n.path[longest:],
				fullPath: n.fullPath,
				handlers: n.handlers,
				priority: n.priority - 1,
				children: n.children,
			}

			n.children = []*node{&child}
			n.path = path[:longest]
			n.fullPath = fullPath[:parentFullPahtIndex+longest]
			n.handlers = nil
		}

		if longest < len(path) {
			path = path[longest:]
		}
	}

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
				n.path = path[:idx]
				path = path[idx:]
			}

			child := node{
				path:     wildcard,
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

	n.path = path
	n.fullPath = fullPath
	n.handlers = handlers
}

type nodeTree struct {
	method string
	root   *node
}

type nodeTrees []nodeTree

func (nts nodeTrees) get(method string) *node {
	for _, tree := range nts {
		if tree.method == method {
			return tree.root
		}
	}

	return nil
}
