package routers

import (
	"golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func Invoice(r *gin.Engine) {
	r.GET("/invoices", controllers.GetInvoices())
	r.GET("/invoices/:invoice_id", controllers.GetInvoice())
	r.POST("/invoices", controllers.CreateInvoice())
	r.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
