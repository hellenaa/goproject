package database

import (
	"../config"
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"os"
)

func ConnectDB() *pgxpool.Pool {
	dbUrl := config.GetKey("dbUrl")
	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
