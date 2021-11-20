package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func readEnv() map[string]string {
	var env map[string]string
	env, err := godotenv.Read("./.env")

	if err != nil {
		log.Fatal("Errore nel caricare il file .env:", err)
	}

	return env
}

func ConnectToDB() *sql.DB {
	var env map[string]string = readEnv()
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		env["DBUSER"], env["DBPASS"], env["DBHOST"], env["DBPORT"], env["DBNAME"])

	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Println("errore: ", err)
	}

	return db
}

func CloseConnectionToDB(db *sql.DB) {
	db.Close()
}
