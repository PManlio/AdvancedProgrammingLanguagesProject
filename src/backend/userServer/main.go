package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./controllers"
	"./models"

	"github.com/gorilla/mux"
)

func homeRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Main Endpoint root /")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	utente := &models.Utente{
		CodFisc:   "asd",
		Nome:      "a",
		Cognome:   "b",
		Email:     "c@d.e",
		Password:  "abcde",
		Citta:     "f",
		Cellulare: "123",
		Genere:    "g",
	}

	json.NewEncoder(w).Encode(utente)

}

func handleRequests() {

	requestHandler := mux.NewRouter().StrictSlash(true)
	pazienteSubrouter := requestHandler.PathPrefix("/paziente").Subrouter()
	// psicologoSubrouter := requestHandler.PathPrefix("/psicologo").Subrouter()

	requestHandler.HandleFunc("/", homeRoot)
	requestHandler.HandleFunc("/getUser", getUser)

	controllers.PazientHandler(pazienteSubrouter)

	// log.Fatal è l'equivalente di Print(), ma seguita da
	// una chiamata a os.Exit(1)
	log.Fatal(http.ListenAndServe(":8085", requestHandler))
}

func main() {
	fmt.Println("Server is running")

	// handleRequests() è "bloccante":
	// le funzioni successive non vengono eseguite
	handleRequests()

	// questa print, dopo handleRequests, non viene eseguita
	// fmt.Println("Server is running")
}
