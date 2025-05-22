package main

import (
	"goproj/view"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.StaticFile("/", "./index.html")
	r.GET("/tasks", view.GetHandler)
	r.POST("/tasks", view.PostHandler)
	r.DELETE("/tasks/:id", view.DeleteHandler)
	r.Run()
}
