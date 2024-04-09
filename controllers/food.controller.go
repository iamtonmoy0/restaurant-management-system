package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iamtonmoy0/restaurant-management-system/database"
	"github.com/iamtonmoy0/restaurant-management-system/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// collections
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

// validator
const validate = validator.New()

// get all foods
func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// get food
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("id")
		var food models.Food
		// find the data
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while finding the data"})
		}
		c.JSON(http.StatusOK, gin.H{
			"data": food,
		})
	}
}

// create food
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		// validation
		validatorErr := validate(food)
		if validatorErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": validatorErr.Error()})
			return
		}
		// finding the menu if exist
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu is not found"})
			return
		}
		// updating the time and price
		food.Updated_at = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		food.Created_at = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		// creating the food
		res, err := foodCollection.InsertOne(ctx, food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
		return

	}

}

// update food
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
func round(num float64, precision int) float64 {

}
