package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Title     string `json:"title"`
	Content   string `json:content`
	CheckMark bool   `json:CheckMark`
	Id        int    `json:id`
}

func GetHandler(c *gin.Context) {
	c.JSON(200, tasks)
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
	{Id: 1, Title: "Покупка хлеба", Content: "Купить 2 буханки хлеба", CheckMark: false},
	{Id: 2, Title: "Покупка яиц", Content: "Купить 2 стака яиц", CheckMark: false},
	{Id: 3, Title: "Покупка альфредычей", Content: "Купить альфредыча", CheckMark: false},
	{Id: 4, Title: "Покупка водки", Content: "Купить водки на пенсию альфредыча", CheckMark: false},
}

func main() {

	r := gin.Default()
	r.StaticFile("/", "./index.html")
	r.GET("/tasks", GetHandler)
	r.POST("/tasks", PostHandler)
	r.Run()
}
