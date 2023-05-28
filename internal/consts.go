package internal

type NodeType uint8

const (
	NodeTypeRoot     NodeType = iota + 1 // Root 路由节点
	NodeTypeParam                        // : 参数路由节点
	NodeTypeCatchAll                     // * 参数路由节点
)
