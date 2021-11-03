package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./models"
)

func homeRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Main Endpoint root /")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	utente := &models.Utente{
		CodFisc:   "asd",
		Nome:      "a",
		Cognome:   "b",
		Email:     "c@d.e",
		Citta:     "f",
		Cellulare: "123",
		Genere:    "g",
	}

	json.NewEncoder(w).Encode(utente)

}

func handleRequests() {
	http.HandleFunc("/", homeRoot)
	http.HandleFunc("/user", getUser)

	// log.Fatal è l'equivalente di Print(), ma seguita da
	// una chiamata a os.Exit(1)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	fmt.Println("Server is running")

	// handleRequests() è "bloccante":
	// le funzioni successive non vengono eseguite
	handleRequests()

	// questa print, dopo handleRequests, non viene eseguita
	// fmt.Println("Server is running")
}
