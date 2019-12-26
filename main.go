package main

import (
	"blog-server/config"
	"blog-server/controller"
	"blog-server/model"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	filename := flag.String("c", "server.yml", "config filename")
	flag.Parse()

	server, db := config.Read(*filename)
	model.Init(db)

	gin.SetMode(server.Mode)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	controller.Init(engine)

	err := engine.Run(fmt.Sprint("0.0.0.0:", server.Port))
	if err != nil {
		panic(err)
	}
}
