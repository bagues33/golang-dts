package main

import (
	"final_project/controllers"
	"final_project/database"
	"final_project/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, err := gorm.Open("sqlite3", "final_project.db")
	if err != nil {
		panic("Failed to connect to the database")
	}
	defer db.Close()

	// Inisialisasi objek database
	database := database.NewDatabase(db)

	r := gin.Default()

	// Grup rute yang memerlukan middleware
	authRequired := r.Group("/")
	authRequired.Use(middlewares.Authenticate())

	// Route Users
	authRequired.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c, database)
	})
	authRequired.PUT("/users/:userId", func(c *gin.Context) {
		controllers.UpdateUser(c, database)
	})
	authRequired.DELETE("/users/:userId", func(c *gin.Context) {
		controllers.DeleteUser(c, database)
	})

	// Route Photos
	authRequired.POST("/photos", func(ctx *gin.Context) {
		controllers.CreatePhoto(ctx, database)
	})
	authRequired.PUT("/photos/:photoId", func(ctx *gin.Context) {
		controllers.UpdatePhoto(ctx, database)
	})
	authRequired.DELETE("/photos/:photoId", func(ctx *gin.Context) {
		controllers.DeletePhoto(ctx, database)
	})

	// Route Comments
	authRequired.POST("/comments", func(c *gin.Context) {
		controllers.CreateComment(c, database)
	})
	authRequired.PUT("/comments/:commentId", func(ctx *gin.Context) {
		controllers.UpdateComment(ctx, database)
	})
	authRequired.DELETE("/comments/:commentId", func(ctx *gin.Context) {
		controllers.DeleteComment(ctx, database)
	})

	// Route Social Media
	authRequired.POST("/social-media", func(ctx *gin.Context) {
		controllers.CreateSocialMedia(ctx, database)
	})
	authRequired.PUT("/social-media/:socialMediaId", func(ctx *gin.Context) {
		controllers.UpdateSocialMedia(ctx, database)
	})
	authRequired.DELETE("/social-media/:socialMediaId", func(ctx *gin.Context) {
		controllers.DeleteSocialMedia(ctx, database)
	})

	// Grup rute yang tidak memerlukan middleware
	public := r.Group("/")
	// public.POST("/users/register", controllers.Register)
	public.POST("/users/register", func(c *gin.Context) {
		controllers.Register(c, database)
	})
	public.GET("/photos", func(c *gin.Context) {
		controllers.GetPhotos(c, database)
	})
	public.GET("/comments", func(ctx *gin.Context) {
		controllers.GetComments(ctx, database)
	})
	public.GET("/social-media", func(ctx *gin.Context) {
		controllers.GetSocialMedias(ctx, database)
	})

	r.Run(":8008")
}
