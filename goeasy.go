package goeasy

import (
	context2 "context"
	"github.com/daobin/goeasy/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type H map[string]any

var once sync.Once
var easy *Engine

// New 新建框架引擎
func New() *Engine {
	once.Do(func() {
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
	})

	return easy
}

// Start 启动框架引擎
func Start(addr string) {
	if easy == nil {
		panic("启动框架引擎失败：框架引擎尚未创建")
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: easy,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(internal.MergeString("启动框架引擎失败：", err.Error()))
	}

	// 等待中断信号以优雅地关闭服务器（设置5秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-quit
	log.Println("服务开始关闭")

	// todo 此处可以做一些关闭前的相关工作，如：资源释放

	ctx, cancel := context2.WithTimeout(context2.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(internal.MergeString("服务关闭失败：", err.Error()))
	}
	log.Println("服务结束关闭")
}
