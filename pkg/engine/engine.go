package engine

import "github.com/gin-gonic/gin"

type Engine struct {
	engine *gin.Engine
}

type Context struct {
	*gin.Context
}

type Handler interface {
	Handle(*Context)
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type HandlerFunc func(*Context) *Response

func New() *Engine {
	return &Engine{
		engine: gin.Default(),
	}
}

func SetMode(mode string) {
	gin.SetMode(mode)
}

func (e *Engine) Run(addr ...string) error {
	return e.engine.Run(addr...)
}

func (h HandlerFunc) handle(c *Context) {
	resp := h(c)
	if resp == nil {
		return
	}
	if resp.Code == 0 {
		c.JSON(200, resp)
	} else {
		c.JSON(500, resp)
	}
}

func (e *Engine) GET(relativePath string, handlers ...HandlerFunc) {
	for _, handler := range handlers {
		e.engine.GET(relativePath, func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
}

func (e *Engine) POST(relativePath string, handlers ...HandlerFunc) {
	for _, handler := range handlers {
		e.engine.POST(relativePath, func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
}

func (e *Engine) Use(handlers ...HandlerFunc) {
	for _, handler := range handlers {
		e.engine.Use(func(c *gin.Context) {
			handler.handle(&Context{c})
		})
	}
}

func (h HandlerFunc) Handle(c *Context) {
	c.JSON(200, h(c))
}
