package controllerls

import "github.com/gin-gonic/gin"

func AddTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
