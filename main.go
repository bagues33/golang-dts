package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// Set up Gin router
	router := gin.Default()
	// Load templates
	router.LoadHTMLGlob("templates/*")

	// Set up route to display status
	router.GET("/status", func(c *gin.Context) {
		status := readStatusFromFile()
		message := getStatusMessage(status.Water, status.Wind)
		c.HTML(http.StatusOK, "status.html", gin.H{
			"status":       status,
			"statusString": message,
		})
	})

	// Start routine to update JSON file every 15 seconds
	go updateStatusPeriodically()

	// Run the server
	router.Run(":8080")
}

// Read status from JSON file
func readStatusFromFile() Status {
	var status Status
	file, err := os.Open("status.json")
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return status
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&status)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return status
	}

	_, err = json.Marshal(&status)
	if err != nil {
		panic(err)
	}

	return status

}

// Write status to JSON file
func writeStatusToFile(status Status) error {
	file, err := os.Create("status.json")
	defer file.Close()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(status)
	if err != nil {
		return err
	}
	return nil
}

// Update status JSON file every 15 seconds with random values
func updateStatusPeriodically() {
	for {
		fmt.Println("Updating status...")
		status := Status{
			Water: rand.Intn(100) + 1,
			Wind:  rand.Intn(100) + 1,
		}
		err := writeStatusToFile(status)
		if err != nil {
			fmt.Println("Error writing status to file:", err)
		}
		time.Sleep(15 * time.Second)
	}
}

// Get status message based on water and wind values
func getStatusMessage(water, wind int) string {
	var statusWater, statusWind string

	if water < 5 {
		statusWater = "Aman"
	} else if water >= 6 && water <= 8 {
		statusWater = "Siaga"
	} else {
		statusWater = "Bahaya"
	}

	if wind < 6 {
		statusWind = "Aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "Siaga"
	} else {
		statusWind = "Bahaya"
	}

	return fmt.Sprintf("Status Air: %s, Status Angin: %s", statusWater, statusWind)
}
