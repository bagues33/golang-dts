package controllers

import (
	"net/http"

	"asignment_2/models"

	"github.com/gin-gonic/gin"

	"asignment_2/database"
)

func GetOrders(c *gin.Context) {
	db := database.GetDB()
	var orders []models.Order

	// Memuat Order beserta Item-nya menggunakan Preload
	db.Preload("Items").Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"message": "All orders",
		"data":    orders,
	})
}

// GetOrder returns a single order
func GetOrder(c *gin.Context) {
	db := database.GetDB()
	order_id := c.Param("id")

	var order models.Order
	if err := db.Preload("Items").Where("id = ?", order_id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order found!",
		"data":    order,
	})

}

// CreateOrder creates a new order
func CreateOrder(c *gin.Context) {
	db := database.GetDB()
	var order models.Order

	// Bind the JSON request body to the Order struct
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the order in the database
	db.Create(&order)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order has been created successfully!",
		"data":    order,
	})
}

// UpdateOrder updates an existing order
func UpdateOrder(c *gin.Context) {
	db := database.GetDB()
	order_id := c.Param("id")

	var order models.Order
	if err := db.Where("id = ?", order_id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Bind the JSON request body to update the order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated order in the database
	db.Save(&order)

	// save the updated item in the database
	for _, item := range order.Items {
		db.Save(&item)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order has been updated successfully!",
		"data":    order,
	})
}

// DeleteOrder deletes an order
func DeleteOrder(c *gin.Context) {
	db := database.GetDB()
	order_id := c.Param("id")

	var order models.Order
	if err := db.Where("id = ?", order_id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.Where("order_id = ?", order_id).Delete(&models.Item{})
	db.Delete(&order)

	c.JSON(http.StatusOK, gin.H{
		"message": "Order has been deleted",
		"data":    true,
	})
}
