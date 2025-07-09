package main

import (
	"golang-restaurant-management/database"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.User(router)
	router.Use(middleware.Authentication())

	router.Food(router)
	router.Menu(router)
	router.Table(router)
	router.Order(router)
	router.OrderItem(router)
	router.Invoice(router)

	router.Run(":" + port)

}
