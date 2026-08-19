package seeders

import (
	"log"

	"gomanagement/database"
	"gomanagement/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	var adminRole models.Role
	var memberRole models.Role

	database.DB.Where("name = ?", "admin").First(&adminRole)
	database.DB.Where("name = ?", "member").First(&memberRole)

	password, err := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)

	if err != nil {
		log.Fatal("Gagal membuat password:", err)
	}

	users := []models.User{
		{
			Name:     "Administrator",
			Email:    "admin@gomanagement.com",
			Password: string(password),
			RoleID:   adminRole.ID,
		},
		{
			Name:     "Member Demo",
			Email:    "member@gomanagement.com",
			Password: string(password),
			RoleID:   memberRole.ID,
		},
	}

	for _, user := range users {
		var existing models.User

		result := database.DB.
			Where("email = ?", user.Email).
			First(&existing)

		if result.Error != nil {
			database.DB.Create(&user)
		}
	}
}