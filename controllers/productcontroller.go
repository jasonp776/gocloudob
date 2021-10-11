package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"psbank.com/gocloudob/database"
	"psbank.com/gocloudob/models"
)

type CreateProductInput struct {
	ID          int       `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       float32   `json:"price" binding:"required"`
	Photo       string    `json:"photo" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreateTime  time.Time `form:"createTime" time_format:"unixNano"`
}

func FindProducts(c *gin.Context) {
	var products []models.Product
	database.Connector.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	product := models.Product{Id: input.ID, Name: input.Name, Category: input.Category, Price: input.Price, Photo: input.Photo, Descriptions: input.Description}
	database.Connector.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /books/:id
// Find a book
func FindProductById(c *gin.Context) {
	var product models.Product

	if err := database.Connector.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// PATCH /books/:id
// Update a book
func UpdateProductByID(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := database.Connector.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Connector.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DELETE /books/:id
// Delete a book
func DeleteProductByID(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := database.Connector.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.Connector.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
