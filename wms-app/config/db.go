package config

import (
	
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
	"database/sql"
)

func GetDBConnection() (*sql.DB, error){
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	var err error
	var conn *sql.DB
	conn, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
		return conn, err
	}
	//defer conn.Close()

	if err = conn.Ping(); err != nil {
		log.Fatal("DB is unreachable:", err)
		return conn, err
	}
	fmt.Println("DB connected successfully")
	return conn, nil
}
