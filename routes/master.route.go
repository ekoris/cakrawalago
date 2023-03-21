package routes

import (
	"cakrawalago/controllers"

	"github.com/gin-gonic/gin"
)

type MasterRouteController struct {
	mailBoxController     controllers.MailBoxController
	bannerPromoController controllers.BannerPromoController
	collateralController  controllers.CollateralController
	bankController        controllers.BankController
	marketController      controllers.MarketController
}

func NewRouteMasterController(marketController controllers.MarketController, bankController controllers.BankController, collateralController controllers.CollateralController, bannerPromoController controllers.BannerPromoController, mailBoxController controllers.MailBoxController) MasterRouteController {
	return MasterRouteController{marketController: marketController, bankController: bankController, collateralController: collateralController, bannerPromoController: bannerPromoController, mailBoxController: mailBoxController}
}

func (m *MasterRouteController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group("master")
	router.GET("/market", m.marketController.Find)
	router.GET("/bank", m.bankController.Find)
	router.GET("/collateral", m.collateralController.Find)
	router.GET("/banner-promo", m.bannerPromoController.Find)
	router.POST("/store-mail-box", m.mailBoxController.Create)
}
