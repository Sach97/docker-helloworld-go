package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	host := "192.168.3.111"
	port := 5432
	user := "postgres-dev"
	password := "s3cr3tp4ssw0rd"
	dbname := "dev"
	sslmode := "disable"

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)

	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error in postgres connection: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
}
