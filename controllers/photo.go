package controllers

import (
	"log"
	"net/http"
	"strconv"

	"final_project/database"
	"final_project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]
	log.Println("userData", userData)

	// Periksa keberadaan pengguna dalam database
	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set userID ke photo
	photo.UserID = int(userID.(float64))

	if err := db.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan data photo yang dibuat dengan user
	if err := db.DB.Preload("User").Find(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

func GetPhotos(c *gin.Context, db *database.Database) {

	var photos []models.Photo
	// show data photos with user

	if err := db.DB.Preload("User").Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func UpdatePhoto(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	var photo models.Photo
	if err := db.DB.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	if photo.UserID != int(userID.(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to update this photo"})
		return
	}

	var inputPhoto models.Photo
	if err := c.ShouldBindJSON(&inputPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&photo).Updates(inputPhoto)

	if err := db.DB.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan data photo yang dibuat dengan user
	if err := db.DB.Preload("User").Find(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)

}

func DeletePhoto(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	var photo models.Photo
	if err := db.DB.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	if photo.UserID != int(userID.(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to delete this photo"})
		return
	}

	db.DB.Delete(&photo)

	c.JSON(http.StatusOK, gin.H{
		"data":    true,
		"message": "photo has been deleted",
	})

}
