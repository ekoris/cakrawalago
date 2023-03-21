package controllers

import (
	"cakrawalago/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BannerPromoController struct {
	DB *gorm.DB
}

func NewBannerPromoController(DB *gorm.DB) BannerPromoController {
	return BannerPromoController{DB}
}

func (bp *BannerPromoController) Find(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var bannerPromos []models.BannerPromo
	results := bp.DB.Limit(intLimit).Offset(offset).Find(&bannerPromos)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(bannerPromos), "data": bannerPromos})
}
