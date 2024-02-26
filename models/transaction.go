package models

import (
	"api/database"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Client struct {
	Saldo  int32 `json:"saldo"`
	Limite int32 `json:"limite"`
}

type Transaction struct {
	ID        int    `json:"id"`
	Valor     int32  `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type TransactionError struct {
	Status  int
	Message string
}

func MakeTransaction(transactionObj *Transaction) (result Client, transactionError *TransactionError) {

	conn := database.GetConnection()

	defer conn.Release()

	sqlTransaction, err := conn.Begin(context.Background())

	if err != nil {
		log.Println("Transaction: Open transaction error")
		panic(err)
	}

	defer sqlTransaction.Rollback(context.Background())

	sqlTransaction.Exec(context.Background(), "SELECT pg_advisory_xact_lock($1)", transactionObj.ID)

	var sqlQuery = `SELECT saldo, limite FROM clientes WHERE id = $1`

	row := sqlTransaction.QueryRow(context.Background(), sqlQuery, transactionObj.ID)

	err = row.Scan(&result.Saldo, &result.Limite)

	if err != nil {

		if err == pgx.ErrNoRows {
			return result, &TransactionError{Status: 404, Message: "User not found"}
		}
		return result, &TransactionError{Status: 500, Message: "Internal server error"}
	}

	if transactionObj.Tipo == "d" && result.Saldo-transactionObj.Valor < result.Limite*-1 {
		return result, &TransactionError{Status: 422, Message: "Limite Indisponível"}
	}

	balanceToBeIncremented := transactionObj.Valor
	if transactionObj.Tipo == "d" {
		balanceToBeIncremented = -transactionObj.Valor
	}

	newClientInformations := sqlTransaction.QueryRow(context.Background(), "UPDATE clientes SET saldo = saldo + $1 WHERE id = $2 RETURNING saldo, limite", balanceToBeIncremented, transactionObj.ID)

	err = newClientInformations.Scan(&result.Saldo, &result.Limite)

	if err != nil {
		log.Println("Transaction error: Error on update client informations")
		panic(err)
	}

	_, err = sqlTransaction.Exec(context.Background(), "INSERT INTO transacoes(valor, tipo,descricao, client_id) VALUES ($1, $2,$3, $4)", transactionObj.Valor, transactionObj.Tipo, transactionObj.Descricao, transactionObj.ID)

	if err != nil {
		log.Println("Transaction error: Error on insert transaction informations")
		panic(err)
	}

	sqlTransaction.Commit(context.Background())

	return result, nil
}
