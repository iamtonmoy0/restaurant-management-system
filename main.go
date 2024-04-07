package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iamtonmoy0/restaurant-management-system/database"
	"github.com/iamtonmoy0/restaurant-management-system/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	r := gin.New()
	r.Use(gin.Logger())
	// middleware

	// router
	routes.UserRoutes(r)
	routes.FoodRoutes(r)
	routes.OrderRoutes(r)
	routes.InvoiceRoutes(r)
	routes.MenuRoutes(r)
	routes.OrderItemRoutes(r)
	routes.TableRoutes(r)

	r.Run(":" + port)
}
