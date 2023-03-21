package controllers

import (
	"cakrawalago/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MailBoxController struct {
	DB *gorm.DB
}

type MailBox struct {
	ID     uint   `gorm:"primaryKey"`
	UserId string `gorm:"not null"`
	Name   string `gorm:"not null;unique"`
	Email  string `gorm:"not null"`
	Value  string `gorm:"not null"`
}

func NewMailBoxController(DB *gorm.DB) MailBoxController {
	return MailBoxController{DB}
}

func (mc *MailBoxController) Create(c *gin.Context) {
	userId := c.PostForm("user_id")
	name := c.PostForm("name")
	email := c.PostForm("email")
	value := c.PostForm("value")

	// Membuat instance dari model
	mailBox := models.MailBox{
		UserId: userId,
		Name:   name,
		Email:  email,
		Value:  value,
	}

	result := mc.DB.Create(&mailBox)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(200, gin.H{"id": mailBox.ID, "name": mailBox.Name, "email": mailBox.Email, "value": mailBox.Value})

}
