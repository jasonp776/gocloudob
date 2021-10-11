package controllers

//POST/*
/*func CreateProfile(c *gin.Context) {
	// Get model if exist
	//	fmt.Println(c.Param("username"))
	var profile models.Profile
	if err := database.Connector.Where("username = ?", c.Param("username")).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Profile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Connector.Model(&profile).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": profile})
}*/
