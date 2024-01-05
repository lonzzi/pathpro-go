package main

import (
	"pathpro-go/conf"
	"pathpro-go/dao"
	"pathpro-go/pkg/engine"
	"pathpro-go/router"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init("config.toml")

	engine.SetMode(gin.DebugMode)
	e := engine.New()
	router.Init(e)
	dao.Init()

	e.Run(":8080")
}
