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
func Start(port string) {
	if easy == nil {
		panic("启动框架引擎失败：框架引擎尚未创建")
	}

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           easy,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(internal.MergeString("启动框架引擎失败：", err.Error()))
	}
}

// Stop 停止框架引擎
func Stop() {
	if easy == nil {
		return
	}

	// todo 待完善
}
