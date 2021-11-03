package main

import (
	"fmt"
	"log"
	"net/http"

	"./models"
)

func homeRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Main Endpoint root /")
}

func handleRequests() {
	http.HandleFunc("/", homeRoot)

	// log.Fatal è l'equivalente di Print(), ma seguita da
	// una chiamata a os.Exit(1)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {

	utente := &models.Utente{
		CodFisc:   "asd",
		Nome:      "a",
		Cognome:   "b",
		Email:     "c@d.e",
		Citta:     "f",
		Cellulare: "123",
		Genere:    "g",
	}

	fmt.Println("esisto", utente)

	handleRequests()

}
