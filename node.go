package goeasy

import "github.com/daobin/goeasy/internal"

type nodeType uint8

type node struct {
	path     string
	fullPath string
	nType    internal.NodeType
	priority uint32
	children []*node
}

func (receiver *node) addRoute(path string, handlers []handlerFunc) {
	fullPath := path
	if receiver.nType == internal.NodeTypeRoot && len(receiver.children) == 0 {
		receiver.addChild(path, fullPath, handlers)
		return
	}
}

func (receiver *node) addChild(path, fullPath string, handlers []handlerFunc) {
	for {
		wildcard, idx, valid := internal.FindPathWildcard(path)
		if idx < 0 {
			break
		}

		if !valid {
			panic(internal.MergeString("路径通配符错误：", fullPath, " >> ", wildcard))
		}
	}
}

type nodeTree struct {
	method string
	root   *node
}

type nodeTrees []nodeTree

func (receiver nodeTrees) get(method string) *node {
	for _, tree := range receiver {
		if tree.method == method {
			return tree.root
		}
	}

	return nil
}
