package controllers

import (
	"cakrawalago/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BankController struct {
	DB *gorm.DB
}

func NewBankController(DB *gorm.DB) BankController {
	return BankController{DB}
}

func (bc *BankController) Find(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var banks []models.AccountBank
	results := bc.DB.Limit(intLimit).Offset(offset).Find(&banks)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(banks), "data": banks})
}
