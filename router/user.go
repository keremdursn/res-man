package routes

import "github.com/gin-gonic/gin"

func User(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers())
	r.GET("/users/:user_id", controllers.GetUser())
	r.POST("/users/signup", controllers.SignUp())
	r.POST("/users/login", controllers.Login())
}
