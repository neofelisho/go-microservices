package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/neofelisho/go-microservices/config"
)

func SayHello(name string) (string, error) {
	conn, err := pgxpool.Connect(context.Background(), config.MustLoad().Database.URI())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	var greeting string
	sqlQuery := fmt.Sprintf("SELECT 'Hello, %s! Greeting from ' || current_database() || ' at ' || current_time", name)
	err = conn.QueryRow(context.Background(), sqlQuery).Scan(&greeting)
	if err != nil {
		return "", err
	}
	return greeting, nil
}
