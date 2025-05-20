package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Title     string `json:"title"`
	CheckMark bool   `json:"checkMark"`
	Id        int    `json:"id"`
}

func GetHandler(c *gin.Context) {
	c.JSON(200, tasks)
}

func DeleteHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, t := range tasks {
		if t.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func PostHandler(c *gin.Context) {
	var userTask Task
	if err := c.BindJSON(&userTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid Json"})
		return
	}
	userTask.Id = len(tasks) + 1
	tasks = append(tasks, userTask)

	c.JSON(http.StatusCreated, userTask)
}

var tasks = []Task{
	{Id: 1, Title: "Покупка хлеба", CheckMark: false},
	{Id: 2, Title: "Покупка яиц", CheckMark: false},
	{Id: 3, Title: "Покупка альфредычей", CheckMark: false},
	{Id: 4, Title: "Покупка водки", CheckMark: false},
}

func main() {

	r := gin.Default()
	r.StaticFile("/", "./index.html")
	r.GET("/tasks", GetHandler)
	r.POST("/tasks", PostHandler)
	r.DELETE("/tasks/:id", DeleteHandler)
	r.Run()
}
