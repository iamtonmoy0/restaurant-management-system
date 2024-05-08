package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamtonmoy0/restaurant-management-system/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

// TODO: need to work on invoice controller 
// get all invoice
func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// get invoice
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// create new invoice
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// update existing invoice
func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
