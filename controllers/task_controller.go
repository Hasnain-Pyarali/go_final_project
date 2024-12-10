package controllers

import (
	"final/config"
	"final/models"
	"net/http"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RenderTasksPage(c *gin.Context) {
    userID := c.Param("user_id")

    var tasks []models.Task
    config.DB.Where("user_id = ?", userID).Find(&tasks)

    c.HTML(200, "tasks.html", gin.H{
        "Tasks": tasks,
        "UserID": userID,
    })
}

func CreateTask(c *gin.Context) {
    userID, _ := strconv.ParseUint(c.PostForm("user_id"), 10, 64)
    title := c.PostForm("title")

    task := models.Task{
        Title:     title,
        Completed: false,
        UserID:    uint(userID),
    }

    config.DB.Create(&task)
    c.Redirect(http.StatusFound, fmt.Sprintf("/tasks/%d", userID))
}

func UpdateTask(c *gin.Context) {
    taskID := c.Param("id")
    title := c.PostForm("title")
    
    var task models.Task
    if err := config.DB.First(&task, taskID).Error; err != nil {
        c.Redirect(http.StatusFound, fmt.Sprintf("/tasks/%d", task.UserID))
        return
    }

    task.Title = title
    config.DB.Save(&task)
    
    c.Redirect(http.StatusFound, fmt.Sprintf("/tasks/%d", task.UserID))
}

func DeleteTask(c *gin.Context) {
    taskID := c.Param("id")
    var task models.Task

    if err := config.DB.First(&task, taskID).Error; err != nil {
        c.Redirect(http.StatusFound, fmt.Sprintf("/tasks/%d", task.UserID))
        return
    }

    config.DB.Delete(&task)
    c.Redirect(http.StatusFound, fmt.Sprintf("/tasks/%d", task.UserID))
}