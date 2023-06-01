package internal

type NodeType uint8

const (
	NodeTypeNormal   NodeType = iota + 1 // 普通路由节点
	NodeTypeParam                        // : 参数路由节点
	NodeTypeCatchAll                     // * 参数路由节点
)
