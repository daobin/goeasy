package internal

type NodeType uint8

const (
	NodeTypeNormal   NodeType = iota // 普通路由节点
	NodeTypeRoot                     // 顶级路由节点
	NodeTypeParam                    // : 参数路由节点
	NodeTypeCatchAll                 // * 参数路由节点
)

const AbortHandlersIndex = 100
