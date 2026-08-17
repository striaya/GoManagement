package controllers 

import (
	"net/http"
	"gomanagement/database"
	"gomanagement/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CreateWorkspaceRequest struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func CreateWorkspace(c *gin.Context) {
	var req CreateWorkspaceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data Workspace tidak valid",
			"error": err.Error(),
		})
		return
	}

	claims, exists := c.Get("claims")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User tidak terautentikasi",
		})
		return
	}

	jwtClaims, ok := claims.(jwt.MapClaims)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Claims tidak valid",
		})
		return
	}

	userIDFloat, ok := jwtClaims["user_id"].(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User ID tidak valid",
		})
		return
	}

	userID := uint(userIDFloat)
	workspace := models.Workspace{
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     userID,
	}

	if err := database.DB.Create(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat workspace",
			"error": err.Error(),
		})
		return
	}

	member := models.WorkspaceMember{
		WorkspaceID: workspace.ID,
		UserID: userID,
		Role: "owner",
	}

	if err := database.DB.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Workspace berhasil dibuat tetapi gagal menambahkan owner",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Workspace berhasil dibuat",
		"workspace": gin.H{
			"id": workspace.ID,
			"name": workspace.Name,
			"description": workspace.Description,
			"owner_id": workspace.OwnerID,
		},
	})
}