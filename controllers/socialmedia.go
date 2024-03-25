package controllers

import (
	"final_project/database"
	"final_project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]

	var socialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate := validator.New()

	// err := validate.Struct(socialMedia)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	socialMedia.UserID = int(userID.(float64))

	if err := db.DB.Create(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Preload("User").Find(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, socialMedia)

}

func GetSocialMedias(c *gin.Context, db *database.Database) {

	var socialMedias []models.SocialMedia
	if err := db.DB.Preload("User").Find(&socialMedias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"social_medias": socialMedias})
}

func UpdateSocialMedia(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	var socialMedia models.SocialMedia

	// validate := validator.New()

	// err := validate.Struct(socialMedia)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	if err := db.DB.Where("id = ?", socialMediaId).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "social media not found"})
		return
	}

	if socialMedia.UserID != int(userID.(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to update this social media"})
		return
	}

	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Preload("User").Find(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func DeleteSocialMedia(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	var socialMedia models.SocialMedia
	if err := db.DB.Where("id = ?", socialMediaId).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "social media not found"})
		return
	}

	if socialMedia.UserID != int(userID.(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to delete this social media"})
		return
	}

	if err := db.DB.Delete(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "social media deleted"})
}
