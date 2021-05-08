package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect()*sql.DB  {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/fines")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
