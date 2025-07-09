package routes

import "github.com/gin-gonic/gin"

func Food(r *gin.Engine) {
	r.GET("/foods", controllers.GetFoods())
	r.GET("/foods/:food_id", controllers.GetFood())
	r.POST("/foods", controllers.CreateFood())
	r.PATCH("/foods/:food_id", controllers.UpdateFood())
}
