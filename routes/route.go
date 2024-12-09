package routes

import (
	"final/controllers"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.CreateTask)

	r.GET("/users/:user_id/tasks", controllers.GetTasksByUser)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)

	return r
}
