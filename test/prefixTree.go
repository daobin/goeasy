package main

import (
	"fmt"
	"strings"
)

type Node struct {
	children map[string]*Node
	isEnd    bool
}
type PrefixTree struct {
	root *Node
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		root: &Node{
			children: make(map[string]*Node),
			isEnd:    false,
		},
	}
}
func (pt *PrefixTree) InsertRoute(route string) {
	routeParts := strings.Split(route, "/")
	currentNode := pt.root
	for _, part := range routeParts {
		if _, ok := currentNode.children[part]; !ok {
			currentNode.children[part] = &Node{
				children: make(map[string]*Node),
				isEnd:    false,
			}
		}
		currentNode = currentNode.children[part]
	}
	currentNode.isEnd = true
}
func (pt *PrefixTree) SearchRoute(route string) bool {
	routeParts := strings.Split(route, "/")
	currentNode := pt.root
	for _, part := range routeParts {
		if _, ok := currentNode.children[part]; !ok {
			return false
		}
		currentNode = currentNode.children[part]
	}
	return currentNode.isEnd
}
func main() {
	prefixTree := NewPrefixTree()
	// 注册路由
	prefixTree.InsertRoute("/home")
	prefixTree.InsertRoute("/home/user")
	prefixTree.InsertRoute("/products")
	prefixTree.InsertRoute("/products/123")
	// 查找路由
	fmt.Println(prefixTree.SearchRoute("/home"))         // 输出: true
	fmt.Println(prefixTree.SearchRoute("/home/user"))    // 输出: true
	fmt.Println(prefixTree.SearchRoute("/home/users"))   // 输出: false
	fmt.Println(prefixTree.SearchRoute("/products"))     // 输出: true
	fmt.Println(prefixTree.SearchRoute("/products/123")) // 输出: true
	fmt.Println(prefixTree.SearchRoute("/products/456")) // 输出: false
	fmt.Println(prefixTree.SearchRoute("/other"))        // 输出: false
}
