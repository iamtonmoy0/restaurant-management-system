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

// update existing Menu
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
