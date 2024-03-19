package main

import (
	"asignment_2/controllers"
	"asignment_2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.StartDB()

	r := gin.Default()

	r.GET("/orders", controllers.GetOrders)
	r.GET("/order/:id", controllers.GetOrder) // router.GET("/person/:id", inDB.GetPerson)
	r.POST("/order", controllers.CreateOrder)
	r.PUT("/order/:id", controllers.UpdateOrder)
	r.DELETE("/order/:id", controllers.DeleteOrder)

	r.Run(":3000")

	// database.StartDB()

}
