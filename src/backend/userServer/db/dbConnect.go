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

/*
func TestSQL() {
	fmt.Println("test sql")

	var env map[string]string = readEnv()

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		env["DBUSER"], env["DBPASS"], env["DBHOST"], env["DBPORT"], env["DBNAME"])

	//connString := fmt.Sprintf("host=%s port=%s user=%s "+
	//	"password=%s dbname=%s",
	//	env["DBHOST"], env["DBPORT"], env["DBUSER"], env["DBPASS"], env["DBNAME"])

	fmt.Println(connString)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Println("errore: ", err)
	}

	defer db.Close()

	type testUtente struct {
		CodFisc   string
		Nome      string
		Cognome   string
		Email     string
		Password  string
		Citta     string
		Cellulare string
		Genere    string
	}

	testusr := testUtente{
		CodFisc:   "codf",
		Nome:      "test",
		Cognome:   "user",
		Email:     "test@user.t",
		Password:  "12345",
		Citta:     "ct",
		Cellulare: "592793",
		Genere:    "fiwjif",
	}

	insert, err := db.Query("INSERT INTO utente(codFisc, nome, cognome, email, password, citta, cellulare, genere) " +
		"VALUES(" + "'" + testusr.CodFisc + "'" + ", " + "'" + testusr.Nome + "'" + ", " + "'" + testusr.Cognome + "'" + ", " + "'" + testusr.Email + "'" + ", " +
		"'" + testusr.Password + "'" + ", " + "'" + testusr.Citta + "'" + ", " + "'" + testusr.Cellulare + "'" + ", " + "'" + testusr.Genere + "'" + ")")
	if err != nil {
		fmt.Println("--------\n\nun errore di query è avvenuto:", err)
	}

	defer insert.Close()

}


func main() {
	// fmt.Println(readEnv())
	TestSQL()
}

*/
