package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gomanagement/database"
	"gomanagement/models"
	"gomanagement/seeders"
	"gomanagement/middleware"
	"gomanagement/controllers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Gagal membaca file .env")
	}

	database.Connect()
	if err := seeders.SeedRoles(database.DB); err != nil {
		panic(err)
	}
	database.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
	)
	router := gin.Default()
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	router.GET(
		"/api/profile",
		middleware.Auth(),
		controllers.Profile,
	)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gomanagement API running",
		})
	})

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "API is healthy",
		})
	})

	router.Run(":8080")
}