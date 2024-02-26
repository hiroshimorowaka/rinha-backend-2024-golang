package controllers

import (
	"api/models"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func TransactionHandler(ctx *gin.Context) {
	clientId := ctx.Param("id")

	ClientIdInt, err := strconv.Atoi(clientId)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "ClientId is not integer",
		})
		return
	}

	var responseBody models.Transaction
	if err := ctx.ShouldBindJSON(&responseBody); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Valor precisa ser inteiro, Tipo e Descrição precisa ser string",
		})
		return
	}

	if err := validatePersonStructure(&responseBody); err != nil {

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	responseBody.ID = ClientIdInt
	result, transactionError := models.MakeTransaction(&responseBody)

	if transactionError != nil {

		ctx.JSON(transactionError.Status, gin.H{"error": transactionError.Message})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func validatePersonStructure(transaction *models.Transaction) error {

	if transaction.Descricao == "" || len(transaction.Descricao) > 10 || len(transaction.Descricao) < 1 {
		return ErrValidation("Descrição")
	}

	if transaction.Valor <= 0 {
		return ErrValidation("Valor")
	}

	reg := regexp.MustCompile(`^[cd]$`)

	if !reg.MatchString(transaction.Tipo) {
		return ErrValidation("Tipo")
	}
	return nil
}

type ErrValidation string

func (e ErrValidation) Error() string {
	return "Invalid Field: " + string(e)
}
