package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gomanagement/database"
	"gomanagement/models"
	"gomanagement/seeders"
	"gomanagement/controllers"
)

func main() {
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