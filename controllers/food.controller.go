package controllers

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iamtonmoy0/restaurant-management-system/database"
	"github.com/iamtonmoy0/restaurant-management-system/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil {
			recordPerPage = 10
		}
		pageNumber, _ := strconv.Atoi(c.Query("page"))
		if err != nil || pageNumber < 1 {
			pageNumber = 1
		}
		startIndex := (pageNumber - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total_count", 1},
					{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
				}}}
		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing food items"})
		}
		var allFoods []bson.M
		if err = result.All(ctx, &allFoods); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allFoods[0])
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
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
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

func round(num float64, precision int) float64 {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// update food
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var menu models.Menu
		var food models.Food
		// get the food id from params
		foodId := c.Param("food_id")
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var updateObj primitive.D
		if food.Name != nil {
			updateObj = append(updateObj, bson.E{"name", food.Name})
		}

		if food.Price != nil {
			updateObj = append(updateObj, bson.E{"price", food.Price})
		}

		if food.Food_image != nil {
			updateObj = append(updateObj, bson.E{"food_image", food.Food_image})
		}

		if food.Menu_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updateObj = append(updateObj, bson.E{"menu", food.Price})
		}
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", food.Updated_at})

		upsert := true
		filter := bson.M{"food_id": foodId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := foodCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprint("foot item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}
