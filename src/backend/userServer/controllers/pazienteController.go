package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	sql "../db"

	"../models"

	"github.com/gorilla/mux"
)

func PazientHandler(pazientRouter *mux.Router) {
	pazientRouter.HandleFunc("/ping", ping)
	pazientRouter.HandleFunc("/create", CreatePazient).Methods("POST")
}

// ---- CRUD Paziente ----

// simple Ping
func ping(w http.ResponseWriter, r *http.Request) {
	ping := struct {
		Ping string
	}{
		Ping: "pong",
	}
	json.NewEncoder(w).Encode(ping)
}

// Crea Paziente:
func CreatePazient(w http.ResponseWriter, r *http.Request) int {
	var paziente models.Paziente

	err := json.NewDecoder(r.Body).Decode(&paziente)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return http.StatusBadRequest
	}

	fmt.Println(paziente)

	db := sql.ConnectToDB()

	insert1, err1 := db.Query("INSERT INTO utente(codFisc, nome, cognome, email, password, citta, cellulare, genere) " +
		"VALUES(" + "'" + paziente.Utente.CodFisc + "'" + ", " + "'" + paziente.Utente.Nome + "'" +
		", " + "'" + paziente.Utente.Cognome + "'" + ", " + "'" + paziente.Utente.Email + "'" + ", " +
		"'" + paziente.Utente.Password + "'" + ", " + "'" + paziente.Utente.Citta + "'" + ", " +
		"'" + paziente.Utente.Cellulare + "'" + ", " + "'" + paziente.Utente.Genere + "');")
	if err1 != nil {
		fmt.Println("--------\n\nun errore di query è avvenuto:", err)
		return http.StatusInternalServerError
	}

	insert2, err2 := db.Query("INSERT INTO paziente(codFisc, patientOf) VALUES(" + "'" + paziente.Utente.CodFisc + "'" + ", " +
		"'" + strings.Join(paziente.PatientOf, ",") + "');")
	if err2 != nil {
		fmt.Println("--------\n\nun errore di query è avvenuto:", err)
		return http.StatusInternalServerError
	}

	defer insert1.Close()
	defer insert2.Close()
	defer sql.CloseConnectionToDB(db)

	return http.StatusOK
}

// Leggi Paziente -
//	- Get Pazient By -
//		- Codice Fiscale:
//		- Mail:
//		- Numero Telefono:

// 	- Get All Pazients:

// Update Paziente
//	- Update Email:

//	- Update num Tel:

//	- Update Genere:

// Rimuovi Paziente:

// ---- CRUD Associazione a Psicologi ----

// aggiungi Psicologo

// leggi Psicologo -
//	- Get Psicologo By CodFisc
//	- Get All Psicologysts

// rimuovi Psicologo
