package controllers

import (
	"net/http"
	"gomanagement/database"
	"gomanagement/models"
	"gomanagement/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak valid",
			"error": err.Error(),
		})
		return
	}

	var existingUser models.User
	result := database.DB.
	Where("email = ?", req.Email).
	First(&existingUser)

	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email sudah digunakan",
		})
		return
	}

	HashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password), bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal!",
		})
		return
	}
	
	var memberRole models.Role
	if err := database.DB.
	Where("name = ?", "Member").
	First(&memberRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Role Member tidak ditemukan",
		})
		return
	}

	user := models.User{
		Name: req.Name,
		Email: req.Email,
		Password: string(HashedPassword),
		RoleID: memberRole.ID,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register berhasil",
		"user": gin.H{
			"id": user.ID,
			"name": user.Name,
			"email": user.Email,
		},
	})
}

func Login(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email atau password wajib diisi",
		})
		return
	}
	var user models.User

	result := database.DB.Preload("Role").
	Where("email = ?", request.Email).
	First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau password salah",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau password salah",
		})
		return
	}
	token, err := utils.GenerateToken(
		user.ID,
		user.RoleID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token": token,
		"user": gin.H{
			"id": user.ID,
			"name": user.Name,
			"email": user.Email,
			"role": user.Role.Name,
		},
	})
}