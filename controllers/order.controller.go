package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamtonmoy0/restaurant-management-system/database"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

// get all Order
func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)

	}
}

// get Order
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// create new Order
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// update existing Order
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
