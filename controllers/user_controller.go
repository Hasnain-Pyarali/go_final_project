package controllers

import (
	"fmt"
	"net/http"
	"final/config"
	"final/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RenderLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func RenderRegisterPage(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func RenderProfilePage(c *gin.Context) {
	var user models.User

	userID := c.Param("user_id")	
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.HTML(200, "profile.html", gin.H{
		"ID": user.ID,
		"Name":  user.Name,
		"Email": user.Email,
		"Password": user.Password,
		"Picture": user.Picture,
	})
}

func Login(c *gin.Context) {
	var user models.User

	email := c.PostForm("email")
	password := c.PostForm("password")
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user.Name})
    c.Redirect(http.StatusFound, fmt.Sprintf("/profile/%d", user.ID))
}

func Register(c *gin.Context) {
	var user models.User

	fullname := c.PostForm("fullname")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	if password != confirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user = models.User{
		Name:     fullname,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

    c.Redirect(http.StatusFound, "/login")
	// c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user.Name})
}

func DeleteAccount(c *gin.Context) {
	userID := c.Param("user_id")

    if err := config.DB.Delete(&models.User{}, userID).Error; err != nil {
        c.JSON(500, gin.H{"error": "Unable to delete account"})
        return
    }

    c.Redirect(http.StatusFound, "/login")
}

// func Delete(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User
	
// 	if err := config.DB.First(&user, id).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, gin.H("error": "User not found"))
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H("error": "Error while searching the user"))
// 		}
// 		return
// 	}
// 	if err := config.DB.Delete(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerErrorgin.H("error": "Error while deleting user account"))
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H("message": "User deleted successfully"))
// }
