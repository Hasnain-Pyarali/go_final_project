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

	r.GET("/tasks/:user_id", controllers.RenderTasksPage)
    r.POST("/tasks", controllers.CreateTask)
    r.POST("/tasks/:id/update", controllers.UpdateTask)
    r.POST("/tasks/:id/delete", controllers.DeleteTask)

	return r
}
