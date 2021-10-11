package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"psbank.com/gocloudob/database"
	"psbank.com/gocloudob/models"
)

type CreateCustomerInput struct {
	ID        int    `json:"id" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Address1  string `json:"address1" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

func FindCustomers(c *gin.Context) {
	var customers []models.Customer
	database.Connector.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// POST /books
// Create new book
func CreateCustomer(c *gin.Context) {
	// Validate input
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	customer := models.Customer{ID: input.ID, FirstName: input.FirstName, LastName: input.LastName, Address1: input.Address1, Email: input.Email}
	database.Connector.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// GET /books/:id
// Find a book
func FindCustomerById(c *gin.Context) {
	var customer models.Customer

	if err := database.Connector.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// PATCH /books/:id
// Update a book
func UpdateCustomerByID(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := database.Connector.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Connector.Model(&customer).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// DELETE /books/:id
// Delete a book
func DeleteCustomerByID(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := database.Connector.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.Connector.Delete(&customer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
