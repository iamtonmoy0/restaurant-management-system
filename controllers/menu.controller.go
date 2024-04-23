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
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// collections
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var validae = validator.New()

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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("id")
		var menu models.Menu
		// find the data
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while finding the data"})
		}
		c.JSON(http.StatusOK, gin.H{
			"data": menu,
		})
	}
}

// create new Menu
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		// validating the struct
		validateErr := validate.Struct(menu)
		if validateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": validateErr.Error()})
			return
		}
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		// creating the menu
		res, err := menuCollection.InsertOne(ctx, menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, res)

	}
}

// in time span
func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}

// update existing Menu
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// menu id
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj primitive.D

		if menu.Start_date != nil && menu.End_date != nil {
			if !inTimeSpan(*menu.Start_date, *menu.End_date, time.Now()) {
				msg := "kindly retype the time"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}
		// update start and end time
		updateObj = append(updateObj, bson.E{"start_date", menu.Start_date})
		updateObj = append(updateObj, bson.E{"end_date", menu.End_date})
		// name
		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{"name": menu.Name})
		}
		// category
		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{"category": menu.Category})
		}
		// time
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at": menu.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := menuCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, opt)
		if err != nil {
			msg := "failed to update "
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
