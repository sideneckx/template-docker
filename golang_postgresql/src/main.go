package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	println("Hi mom")
	host := "127.0.0.1"
	// host := "0.0.0.0"
	// host := "database"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "postgres"

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	// connStr := "user=postgres password=postgres dbname=postgres sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		println("err")
		log.Fatal(err)
	}
	defer db.Close()
	//Create table
	rows, err := db.Query("DROP TABLE IF EXISTS test; CREATE TABLE test (id serial PRIMARY KEY, name varchar(255));")
	if err != nil {
		println("err4")
		log.Fatal(err)
	}
	// columns, err := rows.Columns()
	// if err != nil {
	// 	println("err3")
	// 	log.Fatal(err)
	// }
	println(rows)
}
