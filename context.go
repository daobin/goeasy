package goeasy

import (
	"encoding/json"
	"github.com/daobin/goeasy/internal"
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

func (c *Context) Json(code int, data any) {
	c.Writer.WriteHeader(code)
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	jsonBytes, _ := json.Marshal(data)
	_, _ = c.Writer.Write(jsonBytes)
}
