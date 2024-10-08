package main

import (
	"github.com/MostafaFouad0/Task-Queue-Processor/controllerls"
	"github.com/MostafaFouad0/Task-Queue-Processor/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/task", controllerls.AddTask)
	r.GET("/task/:id", controllerls.GetTask)
	r.Run()
}
