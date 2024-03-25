package controllers

import (
	"net/http"
	"strconv"

	"final_project/database"
	"final_project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserID = int(userID.(float64))

	if err := db.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Preload("User").Preload("Photo").Find(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)

}

func GetComments(c *gin.Context, db *database.Database) {

	var comments []models.Comment
	if err := db.DB.Preload("User").Preload("Photo").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func UpdateComment(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	var comment models.Comment
	if err := db.DB.Where("id = ?", commentId).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	if comment.UserID != int(userID.(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to update this comment"})
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&comment).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Where("id = ?", commentId).First(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Preload("User").Preload("Photo").Find(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)

}

func DeleteComment(c *gin.Context, db *database.Database) {

	users, _ := c.Get("user")
	userData := users.(jwt.MapClaims)
	userID := userData["id"]

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	var comment models.Comment
	if err := db.DB.Where("id = ?", commentId).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	if comment.UserID != int(userID.(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to delete this comment"})
		return
	}

	if err := db.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})

}
