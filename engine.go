package goeasy

import (
	"net/http"
)

// Engine 框架引擎
type Engine struct {
	router
	nodes []*node
}

func (e Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

// 校验是否实现相关接口
var _ http.Handler = (*Engine)(nil)
