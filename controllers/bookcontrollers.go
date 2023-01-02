package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhompra/models"
)

func Create(c *gin.Context) {
	var book models.Books

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	if err := models.DB.Create(&book).Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, book)
}
func GetAll(c *gin.Context) {
	var book []models.Books

	if err := models.DB.Find(&book).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, book)
}
func GetById(c *gin.Context) {
	var book models.Books
	id := c.Param("id")
	if models.DB.Where("id=?", id).Find(&book).RowsAffected == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "data not found",
		})
		return
	}
	c.JSON(200, book)
}
func Delete(c *gin.Context) {
	var book models.Books
	id := c.Param("id")
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "data not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"id":      id,
		"message": "Data deleted",
	})
}
func Put(c *gin.Context) {
	var book models.Books
	id := c.Param("id")

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	if models.DB.Where("id=?", id).Updates(&book).RowsAffected == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "data not found",
		})
		return
	}
	models.DB.Save(&book)
	c.JSON(200, gin.H{
		"message": "Data updated",
	})
}
