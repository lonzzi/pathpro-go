package main

import (
	"pathpro-go/conf"
	"pathpro-go/dao"
	"pathpro-go/pkg/engine"
	"pathpro-go/router"

	_ "pathpro-go/docs"
)

// @title           PathPro API
// @version         1.0
// @description     PathPro API Server.
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.basic  BasicAuth
func main() {
	conf.Init("config.toml")
	e := engine.New()
	router.Init(e)
	dao.Init()

	e.Run(":8080")
}
