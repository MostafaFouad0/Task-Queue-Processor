package main

import (
	"github.com/MostafaFouad0/Task-Queue-Processor/initializers"
	"github.com/MostafaFouad0/Task-Queue-Processor/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Task{})
}
