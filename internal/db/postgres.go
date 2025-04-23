package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() { 
	connStr := "user=bichngoc password=12345678 dbname=bichngoc sslmode=disable"
	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to PostgreSQL")
}