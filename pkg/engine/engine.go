package engine

import (
	"pathpro-go/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	engine *gin.Engine
}

type Context struct {
	*gin.Context
}

type Group struct {
	*gin.RouterGroup
}

type Response struct {
	Code errno.ErrCode `json:"code"`
	Msg  string        `json:"msg"`
	Data interface{}   `json:"data"`
}

type rawResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Handler interface {
	Handle(*Context)
}

type HandlerFunc func(*Context) *Response

type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes

	// StaticFile(string, string) IRoutes
	// StaticFileFS(string, string, http.FileSystem) IRoutes
	// Static(string, string) IRoutes
	// StaticFS(string, http.FileSystem) IRoutes
}

func New() *Engine {
	return &Engine{
		engine: gin.Default(),
	}
}

func SetMode(mode string) {
	gin.SetMode(mode)
}

func (h HandlerFunc) handle(c *Context) {
	resp := h(c)
	if resp == nil {
		return
	}
	c.JSON(resp.Code.HTTPStatusCode, &rawResponse{
		Code: resp.Code.Code,
		Msg:  resp.Msg,
		Data: resp.Data,
	})
	c.Abort()
}

func (e *Engine) Run(addr ...string) error {
	return e.engine.Run(addr...)
}

func (e *Engine) Use(handlers ...HandlerFunc) IRoutes {
	for _, handler := range handlers {
		e.engine.Use(func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
	return e
}

func (e *Engine) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, handler := range handlers {
		e.engine.Handle(httpMethod, relativePath, func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
	return e
}

func (e *Engine) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, handler := range handlers {
		e.engine.GET(relativePath, func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
	return e
}

func (e *Engine) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, handler := range handlers {
		e.engine.POST(relativePath, func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
	return e
}

func (g *Group) Use(handlers ...HandlerFunc) IRoutes {
	for _, handler := range handlers {
		g.RouterGroup.Use(func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
	return g
}

func (g *Group) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, handler := range handlers {
		g.RouterGroup.Handle(httpMethod, relativePath, func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
	return g
}

func (g *Group) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, handler := range handlers {
		g.RouterGroup.GET(relativePath, func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
	return g
}

func (g *Group) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, handler := range handlers {
		g.RouterGroup.POST(relativePath, func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
	return g
}

func (e *Engine) Group(relativePath string, handlers ...HandlerFunc) *Group {
	group := e.engine.Group(relativePath)
	return &Group{group}
}

func (g *Group) Group(relativePath string, handlers ...HandlerFunc) *Group {
	group := g.RouterGroup.Group(relativePath)
	return &Group{group}
}

func (h HandlerFunc) Handle(c *Context) {
	c.JSON(200, h(c))
}
