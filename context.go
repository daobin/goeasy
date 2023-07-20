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

func (c *Context) reset() {
	c.handlersIndex = -1
}

func (c *Context) Next() {
	// 已中止
	if c.handlersIndex == internal.AbortHandlersIndex {
		return
	}

	c.handlersIndex++
	handlerLen := len(c.handlers)

	for c.handlersIndex < handlerLen {
		c.handlers[c.handlersIndex](c)
		c.handlersIndex++
	}
}

func (c *Context) render(code int, r render.IRender) {
	r.Render(code, c.Writer)
}

func (c *Context) Json(code int, data any) {
	c.render(code, &render.Json{Data: data})
}

func (c *Context) bindWith(target any, b binder.IBinder) error {
	return b.Bind(target, c.Request)
}

func (c *Context) BindJson(target any) error {
	return c.bindWith(target, &binder.Json{})
}

func (c *Context) BindQuery(target any) error {
	return c.bindWith(target, &binder.Query{})
}
