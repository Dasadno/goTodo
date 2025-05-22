package view

import (
	"goproj/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	c.JSON(200, model.Tasks)
}

func DeleteHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, t := range model.Tasks {
		if t.Id == id {
			model.Tasks = append(model.Tasks[:i], model.Tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func PostHandler(c *gin.Context) {
	var userTask model.Task
	if err := c.BindJSON(&userTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid Json"})
		return
	}
	userTask.Id = len(model.Tasks) + 1
	model.Tasks = append(model.Tasks, userTask)

	c.JSON(http.StatusCreated, userTask)
}
