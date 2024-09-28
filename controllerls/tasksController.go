package controllerls

import (
	"fmt"
	"os"
	"sync"

	"github.com/MostafaFouad0/Task-Queue-Processor/initializers"
	"github.com/MostafaFouad0/Task-Queue-Processor/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

var taskQueue = make(chan models.Task, 100) // Buffered channel as a task queue
var wg sync.WaitGroup

func init() {
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
}

func AddTask(c *gin.Context) {
	var body struct {
		Status  string
		To      string
		Subject string
		Body    string
	}
	// Get data off the request body
	c.Bind(&body)
	task := models.Task{Status: "Pending", To: body.To, Subject: body.Subject, Body: body.Body}
	result := initializers.DB.Create(&task)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed",
			"error":   result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Task Created",
		"id":      task.ID,
	})
	EnqueueTask(task)
}

func worker(id int) {
	for task := range taskQueue {
		flag := false
		task.Status = "in-progress"
		initializers.DB.Save(&task)

		// Email Details

		m := gomail.NewMessage()
		m.SetHeader("From", os.Getenv("FROM"))
		m.SetHeader("To", task.To)
		m.SetHeader("Subject", task.Subject)
		m.SetBody("text/html", fmt.Sprintf("<p> %v</p>", task.Body))
		d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("FROM"), os.Getenv("PASSWORD"))

		// Retry 3 times before failing

		for i := 0; i < 3; i++ {
			if err := d.DialAndSend(m); err != nil {
				fmt.Printf("Error happend with worker number %v --- error : %v\n", id, err)
			} else {
				flag = true
				break
			}
		}
		if flag {
			task.Status = "OK"
		} else {
			task.Status = "Failed"
		}
		initializers.DB.Save(&task)
		wg.Done()
	}
}

func EnqueueTask(task models.Task) {
	wg.Add(1)
	taskQueue <- task
}
