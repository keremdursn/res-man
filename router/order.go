package routes

import "github.com/gin-gonic/gin"

func Order(r *gin.Engine) {
	r.GET("/orders", controllers.GetOrders())
	r.GET("/orders/:order_id", controllers.GetOrder())
	r.POST("/orders", controllers.CreateOrder())
	r.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
