package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountUserController struct {
	DB *gorm.DB
}

func NewAccountUserController(DB *gorm.DB) AccountUserController {
	return AccountUserController{DB}
}

func (auc *AccountUserController) Create(c *gin.Context) {
	identity, err := uploadFile(c, "identity_attachment")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"path": identity})
}
func uploadFile(c *gin.Context, nameFile string) (string, error) {
	file, err := c.FormFile(nameFile)
	if err != nil {
		return "", err
	}
	filename := file.Filename
	if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
		return "", err
	}
	return "./uploads/" + filename, nil
}
