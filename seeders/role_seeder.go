package seeders

import (
	"gomanagement/database"
	"gomanagement/models"
)

func SeedRoles() {
	roles := []models.Role{
		{Name: "admin"},
		{Name: "project_manager"},
		{Name: "member"},
	}

	for _, role := range roles {
		var existing models.Role

		result := database.DB.
			Where("name = ?", role.Name).
			First(&existing)

		if result.Error != nil {
			database.DB.Create(&role)
		}
	}
}