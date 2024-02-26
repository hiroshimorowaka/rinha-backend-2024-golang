package routes

import (
	controllers "api/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) *gin.RouterGroup {

	api := router.Group("/")
	{
		api.POST("/clientes/:id/transacoes", controllers.TransactionHandler)
		api.GET("/clientes/:id/extrato", controllers.StatementHandler)
	}
	return api
}
