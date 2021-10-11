package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
	"psbank.com/gocloudob/database"
	"psbank.com/gocloudob/models"
)

func FindCredentials(c *gin.Context) {
	var credentials []models.Credentials
	database.Connector.Find(&credentials)

	c.JSON(http.StatusOK, gin.H{"data": credentials})
}

func Signup(c *gin.Context) {

	var input models.Credentials
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	credential := models.Credentials{Username: input.Username, Email: input.Email, Password: input.Password}
	if err := database.Connector.Where("username = ?", input.Username).First(&credential).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exist"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err == nil {

		creds := models.Credentials{Password: string(hashedPassword), Username: credential.Username, Email: credential.Email}
		database.Connector.Create(&creds)
		c.JSON(http.StatusOK, gin.H{"success": creds})

	}

}

func Login(c *gin.Context) { // Get model if exist

	var input models.Credentials
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	credential := models.Credentials{Username: input.Username, Password: input.Password}
	if err := database.Connector.Where("username = ?", credential.Username).First(&credential).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(credential.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"data": "Invalid Password"})
		return
	}
	/*session := sessions.Default(c)
	//sessionid := shortuuid.New()
	//session.Set("id", 12090292)
	/session.Set("email", "jpascual@gmail.com")
	session.Save() */
	c.JSON(http.StatusOK, gin.H{"data": "User LoggedIn"})
	return

}

func Login2(c *gin.Context) {

	session := sessions.Default(c)
	session.Set("id", 12090292)
	session.Set("email", "test@gmail.com")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully",
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}
