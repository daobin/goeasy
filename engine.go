package goeasy

import (
	"net/http"
	"strings"
	"sync"
)

// Engine 框架引擎
type Engine struct {
	router
	ctxPool sync.Pool
	trees   nodeTrees
}

// newContext 新建上下文
func (e *Engine) newContext() *Context {
	return &Context{}
}

// addRouteNode 添加路由节点（注册路由）
func (e *Engine) addRouteNode(httpMethod, absolutePath string, handlers []handlerFunc) {
	root := e.trees.get(httpMethod)
	if root == nil {
		root = &node{
			fullPath: "/",
			children: map[string]*node{},
		}
		e.trees = append(e.trees, nodeTree{
			method: httpMethod,
			root:   root,
		})
	}

	root.addRouteNode(absolutePath, handlers)
}

// ServeHTTP 实现HTTP服务（监听路由）
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := e.ctxPool.Get().(*Context)
	c.Writer = w
	c.Request = req
	c.reset()

	e.handleHttpRequest(c)
	e.ctxPool.Put(c)
}

func (e *Engine) handleHttpRequest(c *Context) {
	req := c.Request
	segments := strings.Split(req.URL.Path, "/")

	n := e.trees.get(req.Method)
	if n == nil {
		c.Json(http.StatusNotFound, H{"status": http.StatusNotFound, "msg": "请求资源不存在"})
		return
	}

	c.Params = map[string]string{}

	for idx, segment := range segments {
		if segment == "" {
			continue
		}

		// 精确匹配优先
		child := n.getChildNode(segment)
		if child == nil {
			c.Json(http.StatusNotFound, H{"status": http.StatusNotFound, "msg": "请求资源不存在"})
			return
		}
		n = child

		// 参数赋值
		if n.path[0] == ':' {
			c.Params[n.path[1:]] = segment[1:]
		} else if n.path[0] == '*' {
			segments[idx] = segment[1:]
			c.Params[n.path[1:]] = strings.Join(segments[idx:], "/")
			break
		}
	}

	if len(n.handlers) == 0 {
		c.Json(http.StatusNotFound, H{"status": http.StatusNotFound, "msg": "请求资源不存在"})
		return
	}

	c.handlers = n.handlers
	c.Next()
}

// 校验是否实现相关接口
var _ http.Handler = (*Engine)(nil)
