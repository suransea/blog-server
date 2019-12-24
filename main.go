package main

import (
	"blog-server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	controller.Setup(engine)

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
