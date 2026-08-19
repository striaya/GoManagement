package seeders

import (
	"gomanagement/database"
	"gomanagement/models"
)

func SeedPriorities() {
	priorities := []models.Priority{
		{
			Name:  "Low",
			Level: 1,
		},
		{
			Name:  "Medium",
			Level: 2,
		},
		{
			Name:  "High",
			Level: 3,
		},
		{
			Name:  "Urgent",
			Level: 4,
		},
	}

	for _, priority := range priorities {
		var existing models.Priority

		result := database.DB.
			Where("name = ?", priority.Name).
			First(&existing)

		if result.Error != nil {
			database.DB.Create(&priority)
		}
	}
}