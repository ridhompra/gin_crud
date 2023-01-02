package controllers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/ridhompra/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9.%+\-]+\.[a-z0-9.%+\-]`)
	return Re.MatchString(email)
}
func SignUp(c *gin.Context) {
	var user models.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if len(string(user.Password)) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password must min 6 character",
		})
		return
	}
	if !ValidateEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email not valid",
		})
		return
	}
	hashedpass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedpass)
	models.DB.Where("email=?", user.Email).Find(&user)
	if user.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exist",
		})
		return
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"Email":     user.Email,
		"username":  user.Username,
		"message":   "Sign up successfully",
		"user_test": user,
	})
}
