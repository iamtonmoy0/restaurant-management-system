package helpers

import (
	"github.com/iamtonmoy0/restaurant-management-system/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var user *mongo.Collection = database.OpenCollection(database.Client, "user")
