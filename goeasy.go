package goeasy

import (
	"fmt"
	"net/http"
)

var easy *Engine

// New 新建框架引擎
func New() *Engine {
	if easy != nil {
		return easy
	}

	easy = &Engine{
		router: router{
			basePath: "/",
		},
	}
	easy.router.engine = easy

	return easy
}

// Start 启动框架引擎
func Start(addr string) {
	if easy == nil {
		panic("启动框架引擎失败：框架引擎尚未创建")
	}

	err := http.ListenAndServe(addr, easy)
	if err != nil {
		panic(fmt.Sprintf("启动框架引擎失败：%s", err.Error()))
	}
}

// Stop 停止框架引擎
func Stop() {
	if easy == nil {
		return
	}

}
