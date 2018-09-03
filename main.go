package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres-dev"
	password = "s3cr3tp4ssw0rd"
	dbname   = "dev"
)

func main() {
	host := "localhost"
	port := "5432"
	user := os.GetEnv("POSTGRES_USER")
	password := os.GetEnv("POSTGRES_PASSWORD")
	dbname := os.GetEnv("POSTGRES_DB")
	sslmode := "disable"

	// host := "localhost"
	// port := 5432
	// user := "postgres-dev"
	// password := "s3cr3tp4ssw0rd"
	// dbname := "dev"
	// sslmode := "disable"

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)

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
