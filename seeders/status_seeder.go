package seeders

import (
	"gomanagement/database"
	"gomanagement/models"
)

func SeedStatuses() {
	statuses := []models.Status{
		{
			Name:     "Todo",
			Color:    "#6B7280",
			Position: 1,
		},
		{
			Name:     "In Progress",
			Color:    "#3B82F6",
			Position: 2,
		},
		{
			Name:     "Review",
			Color:    "#F59E0B",
			Position: 3,
		},
		{
			Name:     "Done",
			Color:    "#10B981",
			Position: 4,
		},
	}

	for _, status := range statuses {
		var existing models.Status

		result := database.DB.
			Where("name = ?", status.Name).
			First(&existing)

		if result.Error != nil {
			database.DB.Create(&status)
		}
	}
}