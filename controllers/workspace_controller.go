package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gomanagement/database"
	"gomanagement/models"
	"net/http"
)

type CreateWorkspaceRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func CreateWorkspace(c *gin.Context) {
	var req CreateWorkspaceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data Workspace tidak valid",
			"error":   err.Error(),
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
			"error":   err.Error(),
		})
		return
	}

	member := models.WorkspaceMember{
		WorkspaceID: workspace.ID,
		UserID:      userID,
		Role:        "owner",
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
			"id":          workspace.ID,
			"name":        workspace.Name,
			"description": workspace.Description,
			"owner_id":    workspace.OwnerID,
		},
	})
}

func GetWorkspaces(c *gin.Context) {
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

	var workspaces []models.Workspace
	if err := database.DB.
		Where("owner_id = ?", userID).
		Find(&workspaces).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil workspace",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Workspace berhasil diambil",
		"workspace": workspaces,
	})
}

func GetWorkspace(c *gin.Context) {
	id := c.Param("id")

	var workspace models.Workspace

	if err := database.DB.First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Workspace tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Workspace berhasil ditemukan",
		"workspace": workspace,
	})
}

type UpdateWorkspaceRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func UpdateWorkspace(c *gin.Context) {
	id := c.Param("id")

	var workspace models.Workspace

	if err := database.DB.First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Workspace tidak ditemukan",
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
		c.JSON(http.StatusAccepted, gin.H{
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

	if workspace.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Kamu tidak memiliki akses ke workspace ini",
		})
		return
	}

	var req UpdateWorkspaceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data workspace tidak valid",
			"error": err.Error(),
		})
		return
	}

	workspace.Name = req.Name
	workspace.Description = req.Description

	if err := database.DB.Save(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal memperbarui workspace",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Workspace berhasil diperbarui",
		"workspace": workspace,
	})
}