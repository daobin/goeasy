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
}

// addRouteNode 添加路由节点（注册路由）
func (n *node) addRouteNode(fullPath string, handlers []handlerFunc) {
	path := fullPath
	n.priority++

	if n.nType == internal.NodeTypeRoot && len(n.handlers) == 0 {
		n.insertWildChild(path, fullPath, handlers)
		return
	}

	parentFullPathIndex := 0

walk:
	for {
		// 公共前缀最大长度
		longest := internal.CommonPrefixLongest(path, n.path)

		// 分裂当前节点，将分裂的后半部分添加为子节点
		if longest < len(n.path) {
			newNode := &node{
				fullPath: n.fullPath,
				path:     n.path[longest:],
				indices:  n.indices,
				handlers: n.handlers,
				nType:    internal.NodeTypeNormal,
				priority: n.priority - 1,
				children: n.children,
			}

			n.fullPath = fullPath[:parentFullPathIndex+longest]
			n.indices = string(n.path[longest])
			n.path = path[:longest]
			n.handlers = nil
			n.children = []*node{newNode}
		}

		// 添加新节点
		if longest < len(path) {
			path = path[longest:]
			pathChar := path[0]

			// 校验path是否存在公共前缀的子节点
			for i, max := 0, len(n.indices); i < max; i++ {
				if pathChar == n.indices[i] {
					parentFullPathIndex += len(n.path)
					i = n.incrementChildPriority(i)
					n = n.children[i]
					continue walk
				}
			}

			newNode := &node{
				fullPath: fullPath,
				nType:    internal.NodeTypeNormal,
			}
			n.addChild(newNode)

			// 优化子节点排序
			n.indices += string([]byte{pathChar})
			n.incrementChildPriority(len(n.indices) - 1)

			n = newNode
			n.insertWildChild(path, fullPath, handlers)
			return
		}

		n.fullPath = fullPath
		n.handlers = handlers
		return
	}
}

// insertWildChild 插入通配符参数子节点
func (n *node) insertWildChild(path, fullPath string, handlers []handlerFunc) {
	//for {
	//	wildcard, idx, valid := internal.FindPathWildcard(path)
	//	if idx < 0 {
	//		break
	//	}
	//
	//	if !valid {
	//		panic(internal.MergeString("路径通配符错误：", fullPath, " >> ", wildcard))
	//	}
	//
	//	if wildcard[0] == ':' {
	//		if idx > 0 {
	//			path = path[idx:]
	//		}
	//
	//		newNode := node{
	//			fullPath: fullPath,
	//			nType:    internal.NodeTypeParam,
	//		}
	//
	//		n.children = append(n.children, &newNode)
	//		n = &newNode
	//		n.priority++
	//
	//		if len(wildcard) < len(path) {
	//			path = path[len(wildcard):]
	//			newNode = node{fullPath: fullPath, priority: 1}
	//			n.children = append(n.children, &newNode)
	//			n = &newNode
	//			continue
	//		}
	//
	//		n.handlers = handlers
	//		return
	//	}
	//}

	// 没有找到通配符时处理
	n.path = path
	n.fullPath = fullPath
	n.handlers = handlers
}

// addChild 添加子节点
func (n *node) addChild(child *node) {
	if n.isWildcard() && len(n.children) > 0 {
		wildcardChild := n.children[len(n.children)-1]
		n.children = append(n.children[:len(n.children)-1], child, wildcardChild)
		return
	}

	n.children = append(n.children, child)
}

// isWildcard 是否为通配符参数节点
func (n *node) isWildcard() bool {
	return n.nType == internal.NodeTypeParam || n.nType == internal.NodeTypeCatchAll
}

// incrementChildPriority 提升子节点priority，并根据子节点priority重排子节点顺序
func (n *node) incrementChildPriority(index int) int {
	n.children[index].priority++
	priority := n.children[index].priority

	newIndex := index
	// 重排子节点顺序
	for ; newIndex > 0 && n.children[newIndex-1].priority < priority; newIndex-- {
		n.children[newIndex-1], n.children[newIndex] = n.children[newIndex], n.children[newIndex-1]
	}

	// 根据子节点顺序重组节点indices
	if newIndex != index {
		n.indices = n.indices[:newIndex] + n.indices[index:index+1] + n.indices[newIndex:index] + n.indices[index+1:]
	}

	return newIndex
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
