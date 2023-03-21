package routes

import (
	"cakrawalago/controllers"

	"github.com/gin-gonic/gin"
)

type AccountUserRouteController struct {
	accountUserController controllers.AccountUserController
}

func NewRouteAccountUserController(accountUserController controllers.AccountUserController) AccountUserRouteController {
	return AccountUserRouteController{accountUserController: accountUserController}
}

func (a *AccountUserRouteController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group("account")
	router.POST("/save", a.accountUserController.Create)
}
