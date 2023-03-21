package controllers

import (
	"cakrawalago/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CollateralController struct {
	DB *gorm.DB
}

func NewCollateralController(DB *gorm.DB) CollateralController {
	return CollateralController{DB}
}

func (cc *CollateralController) Find(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var collaterals []models.Collateral
	results := cc.DB.Limit(intLimit).Offset(offset).Find(&collaterals)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(collaterals), "data": collaterals})
}
