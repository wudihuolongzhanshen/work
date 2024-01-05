package main

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.CreateTodo)
	}
	r.Run(":8080")
}
