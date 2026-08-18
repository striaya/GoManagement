package controllers 

import (
	"net/http"
	"gomanagement/database"
	"gomanagement/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CreateProjectRequest struct {
	WorkspaceID uint `json:"workspace_id", binding:"required"`
	Name string `json:"name", binding:"required"`
	Key string `json:"key", binding:"required"`
	Description string `json:"description"`
}

func CreateProject(c *gin.Context) {
	var req CreateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data project tidak valid",
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

jwtClaims := claims.(jwt.MapClaims)

userIDFloat, ok := jwtClaims["user_id"].(float64)

if !ok {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "User ID tidak valid",
	})
	return
}

userID := uint(userIDFloat)

	var workspace models.Workspace

	if err := database.DB.Where("id = ?", req.WorkspaceID).First(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Workspace tidak ditemukan",
		})
		return
	}

	var member models.WorkspaceMember

	if err := database.DB.
	Where("workspace_id = ? AND user_id = ?", req.WorkspaceID, userID).
	First(&member).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Kamu bukan member workspace ini",
		})
		return
	}

		project := models.Project{
		WorkspaceID: req.WorkspaceID,
		OwnerID:     userID,
		Name:        req.Name,
		Key:         req.Key,
		Description: req.Description,
		Status:      "active",
	}

	if err := database.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat project",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Project berhasil dibuat",
		"project": gin.H{
			"id":           project.ID,
			"workspace_id": project.WorkspaceID,
			"owner_id":     project.OwnerID,
			"name":         project.Name,
			"key":          project.Key,
			"description":  project.Description,
			"status":       project.Status,
		},
	})
}