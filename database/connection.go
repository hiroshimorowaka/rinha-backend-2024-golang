package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitConnectionPool() error {

	config, err := pgxpool.ParseConfig("postgres://docker:docker@localhost:5432/rinha?sslmode=disable")

	if err != nil {
		log.Println(err.Error())
	}

	config.MaxConnIdleTime = 0

	config.MaxConns = 16 // 16
	config.MinConns = 13 // 13

	// config.HealthCheckPeriod = time.Second * 10

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	log.Println(pool.Config())

	db = pool
	return nil
}

func GetConnection() *pgxpool.Conn {

	conn, err := db.Acquire(context.Background())

	if err != nil {
		log.Println("Error when connecting to pool")
		panic(err)
	}

	return conn
}
