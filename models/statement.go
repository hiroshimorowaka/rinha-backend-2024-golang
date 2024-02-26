package models

import (
	"api/database"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Extrato struct {
	Valor       int32     `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type ClientInformations struct {
	Total       int32     `json:"total"`
	Limite      int32     `json:"limite"`
	DataExtrato time.Time `json:"data_extrato"`
}

type ClientStatement struct {
	Informations      ClientInformations
	UltimasTransacoes []Extrato `json:"ultimas_transacoes"`
}

type StatementError struct {
	Status  int
	Message string
}

func GetStatement(ClientId int) (ultimasTransacoes ClientStatement, statementError *StatementError) { //começar a retornar erros aqui pra tratar depois

	conn := database.GetConnection()

	defer conn.Release()

	sqlTransaction, err := conn.Begin(context.Background())

	if err != nil {
		log.Println("Statement: Open Transaction error")
		panic(err)
	}

	defer sqlTransaction.Rollback(context.Background()) // Rollback da transação se ocorrer um erro ou não for confirmada

	var clientQuery = `SELECT saldo, limite FROM clientes WHERE id = $1`

	row := sqlTransaction.QueryRow(context.Background(), clientQuery, ClientId)

	err = row.Scan(&ultimasTransacoes.Informations.Total, &ultimasTransacoes.Informations.Limite)

	if err != nil {
		if err == pgx.ErrNoRows {
			return ultimasTransacoes, &StatementError{Status: 404, Message: "User not found"}
		}
		return ultimasTransacoes, &StatementError{Status: 500, Message: "Internal server error"}
	}

	var databaseQuery = `
	(select saldo as valor, 'valor' as tipo, 'valor' as descricao, now() as realizada_em from clientes where id = $1) 
	union all 
	(select valor, tipo, descricao, realizada_em from transacoes
	where client_id = $1
	order by id desc limit 10)`

	rows, err := sqlTransaction.Query(context.Background(), databaseQuery, ClientId)

	if err != nil {
		log.Println("Statement: Query recent transactions error")
		panic(err)
	}

	for rows.Next() {

		var valor int32
		var tipo string
		var descricao string
		var realizadaEm time.Time

		err := rows.Scan(&valor, &tipo, &descricao, &realizadaEm)
		if err != nil {
			panic(err)
		}

		if tipo == "valor" {
			ultimasTransacoes.Informations.DataExtrato = realizadaEm
		} else {
			ultimasTransacoes.UltimasTransacoes = append(ultimasTransacoes.UltimasTransacoes, Extrato{
				Valor:       valor,
				Tipo:        tipo,
				Descricao:   descricao,
				RealizadaEm: realizadaEm,
			})
		}
	}

	if err = rows.Err(); err != nil {
		log.Println("Statement: Scan recent transactions error")
		panic(err)
	}

	return

}
