package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect()*sql.DB  {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/fines")
	//defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Mysql",version," connected")

	return db
}
