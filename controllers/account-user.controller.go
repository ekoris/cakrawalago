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
	identity := c.FormFile("identity_attachment")
	// path, err := uploadFileAndGetPath(c, identity)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"path": identity})
}

// func uploadFileAndGetPath(c *gin.Context, file) (string, error) {
// 	err := c.Request.ParseMultipartForm(32 << 20) // 32 MB
// 	if err != nil {
// 		return "", err
// 	}
// 	file, header, err := file
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	// Create a file on the server to write to
// 	out, err := os.Create("uploads/" + header.Filename)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer out.Close()

// 	// Write the file contents to disk
// 	_, err = io.Copy(out, file)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Return success response
// 	return out.Name(), nil
// }
