package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamtonmoy0/restaurant-management-system/database"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// collections
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

// get all Menu
func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(ctx, bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var allMenus []bson.M
		if err := result.All(ctx, &allMenus); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

// get Menu
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// create new Menu
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// update existing Menu
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
