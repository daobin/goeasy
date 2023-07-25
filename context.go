package goeasy

import (
	"github.com/daobin/goeasy/internal"
	"github.com/daobin/goeasy/internal/binder"
	"github.com/daobin/goeasy/internal/render"
	"net/http"
)

type Context struct {
	Writer        http.ResponseWriter
	Request       *http.Request
	handlers      handlerChain
	handlersIndex int
}

// reset 重置相关数据
func (c *Context) reset() {
	c.handlersIndex = -1
}

// Abort 中断请求
func (c *Context) Abort() {
	c.handlersIndex = internal.AbortHandlersIndex
}

// Next 执行下一个handler
func (c *Context) Next() {
	c.handlersIndex++
	handlerLen := len(c.handlers)

	for c.handlersIndex < handlerLen {
		c.handlers[c.handlersIndex](c)
		c.handlersIndex++
	}
}

// render 渲染数据
func (c *Context) render(code int, r render.IRender) {
	r.Render(code, c.Writer)
}

// Json 渲染JSON数据
func (c *Context) Json(code int, data any) {
	c.render(code, &render.Json{Data: data})
}

// bindWith 绑定请求数据到指定对象
func (c *Context) bindWith(target any, b binder.IBinder) error {
	return b.Bind(target, c.Request)
}

// BindJson 绑定JSON数据到指定对象
func (c *Context) BindJson(target any) error {
	return c.bindWith(target, &binder.Json{})
}

// BindQuery 绑定Query数据到指定对象
func (c *Context) BindQuery(target any) error {
	return c.bindWith(target, &binder.Query{})
}
