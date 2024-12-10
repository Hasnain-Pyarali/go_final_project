package routes

import (
	"final/controllers"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/login")
    })

	r.GET("/login", controllers.RenderLoginPage)
	r.GET("/register", controllers.RenderRegisterPage)
	r.GET("/profile/:user_id", controllers.RenderProfilePage)
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.POST("/update-profile-picture", controllers.UpdateProfilePicture)
	r.POST("/delete-account/:user_id", controllers.DeleteAccount)

	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.CreateTask)

	r.GET("/users/:user_id/tasks", controllers.GetTasksByUser)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)

	return r
}
