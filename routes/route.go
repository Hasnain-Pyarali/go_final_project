package routes

import (
	"github.com/gin-gonic/gin"
	"final/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.CreateTask)

	return r
}
