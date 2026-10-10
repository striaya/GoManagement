func DeleteWorkspace(c *gin.Context) {
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

	// Hanya owner yang boleh menghapus workspace
	if workspace.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Kamu tidak memiliki akses untuk menghapus workspace ini",
		})
		return
	}

	if err := database.DB.Delete(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus workspace",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Workspace berhasil dihapus",
	})
}