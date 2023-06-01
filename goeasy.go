package goeasy

import (
	"github.com/daobin/goeasy/internal"
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
			basePath:    "/",
			handlers:    nil,
			engine:      nil,
			isEngineNew: true,
		},
	}
	easy.router.engine = easy
	easy.ctxPool.New = func() any {
		return easy.newContext()
	}

	return easy
}

// Start 启动框架引擎
func Start(addr string) {
	if easy == nil {
		panic("启动框架引擎失败：框架引擎尚未创建")
	}

	err := http.ListenAndServe(addr, easy)
	if err != nil {
		panic(internal.MergeString("启动框架引擎失败：", err.Error()))
	}
}

// Stop 停止框架引擎
func Stop() {
	if easy == nil {
		return
	}

}
