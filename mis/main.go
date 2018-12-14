package main

import (
	"./handler"
	"./middleware"
	"github.com/gin-gonic/gin"
)

func register_handler(r *gin.Engine) {
	r.GET("/", handler.HomeHandler)
	r.POST("/add", handler.AddRecordHandler)
	r.GET("/del", handler.DelRecordHandler)
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("./view/*")
	r.Static("static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.Use(middleware.Database("./database.sqlite3"))

	register_handler(r)

	r.Run(":80")
}
