package controllers

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StatementHandler(ctx *gin.Context) {
	clientId := ctx.Param("id")

	clientIdInt, err := strconv.Atoi(clientId)

	if err != nil {

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "ClientId is not integer",
		})
		return
	}

	result, statementErr := models.GetStatement(clientIdInt)

	if statementErr != nil {
		ctx.JSON(statementErr.Status, gin.H{"error": statementErr.Message})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"saldo": map[string]interface{}{
			"total":        result.Informations.Total,
			"data_extrato": result.Informations.DataExtrato,
			"limite":       result.Informations.Limite,
		},
		"ultimas_transacoes": result.UltimasTransacoes,
	})
}
