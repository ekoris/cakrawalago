package main

import (
	"cakrawalago/configs"
	"cakrawalago/controllers"
	"cakrawalago/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	AccountUserController      controllers.AccountUserController
	AccountUserRouteController routes.AccountUserRouteController

	MailBoxController     controllers.MailBoxController
	BannerPromoController controllers.BannerPromoController
	CollateralController  controllers.CollateralController
	BankController        controllers.BankController
	MarketController      controllers.MarketController
	MasterRouteController routes.MasterRouteController
)

func init() {
	configs.SetupDB()

	MailBoxController = controllers.NewMailBoxController(configs.DB)
	BannerPromoController = controllers.NewBannerPromoController(configs.DB)
	CollateralController = controllers.NewCollateralController(configs.DB)
	BankController = controllers.NewBankController(configs.DB)
	MarketController = controllers.NewMarketController(configs.DB)
	MasterRouteController = routes.NewRouteMasterController(MarketController, BankController, CollateralController, BannerPromoController, MailBoxController)

	AccountUserController = controllers.NewAccountUserController(configs.DB)
	AccountUserRouteController = routes.NewRouteAccountUserController(AccountUserController)

	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000"}
	server.Use(cors.New(corsConfig))
	router := server.Group("/api")
	MasterRouteController.SetupRoutes(router)
	AccountUserRouteController.SetupRoutes(router)
	log.Fatal(server.Run(":8000"))
}
