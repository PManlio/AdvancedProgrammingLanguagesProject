package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	myDBpckg "../db"
	"../models"
	"../utils"

	"github.com/gorilla/mux"
)

func PazientHandler(pazientRouter *mux.Router) {
	// Simple Ping
	pazientRouter.HandleFunc("/ping", ping)

	// CRUD paziente:
	pazientRouter.HandleFunc("/create", CreatePazient).Methods("POST")
	pazientRouter.HandleFunc("/getbycodfisc", getPazientByCodFisc).Methods("GET")
	pazientRouter.HandleFunc("/getbyemail", getPazientByMail).Methods("GET")
	pazientRouter.HandleFunc("/getbyphonenumber", getPazientByPhoneNumber).Methods("GET")
	pazientRouter.HandleFunc("/getallpatients", getAllPatients).Methods("GET")
}

// simple Ping
func ping(w http.ResponseWriter, r *http.Request) {
	ping := struct {
		Ping string
	}{
		Ping: "pong",
	}
	json.NewEncoder(w).Encode(ping)
}

// ---- CRUD Paziente ----

// Crea Paziente:
func CreatePazient(w http.ResponseWriter, r *http.Request) {

	var paziente models.Paziente
	err := json.NewDecoder(r.Body).Decode(&paziente)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Println(paziente)
	paziente.Utente.Password = utils.Encrypt(paziente.Utente.Password)
	// fmt.Println(paziente)

	db := myDBpckg.ConnectToDB()

	// CONTROLLO SU TABELLA PER VERIFICARE ESISTENZA
	var isPresent bool

	errCheck := db.QueryRow("SELECT codFisc FROM utente INNER JOIN paziente ON utente.codFisc = paziente.codFisc " +
		"WHERE paziente.codFisc = '" + paziente.Utente.CodFisc + "'").Scan(&isPresent)

	if errCheck != nil && errCheck != sql.ErrNoRows {

		http.Error(w, "row exists already", http.StatusForbidden)

	} else if !isPresent {

		insert1, err1 := db.Query("INSERT INTO utente(codFisc, nome, cognome, email, password, citta, cellulare, genere) " +
			"VALUES(" + "'" + paziente.Utente.CodFisc + "'" + ", " + "'" + paziente.Utente.Nome + "'" +
			", " + "'" + paziente.Utente.Cognome + "'" + ", " + "'" + paziente.Utente.Email + "'" + ", " +
			"'" + paziente.Utente.Password + "'" + ", " + "'" + paziente.Utente.Citta + "'" + ", " +
			"'" + paziente.Utente.Cellulare + "'" + ", " + "'" + paziente.Utente.Genere + "');")
		if err1 != nil {
			fmt.Println("--------\n\nun errore di query è avvenuto:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		insert2, err2 := db.Query("INSERT INTO paziente(codFisc, patientOf) VALUES(" + "'" + paziente.Utente.CodFisc + "'" + ", " +
			"'" + strings.Join(paziente.PatientOf, ",") + "');")
		if err2 != nil {
			fmt.Println("--------\n\nun errore di query è avvenuto:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer insert1.Close()
		defer insert2.Close()

		json.NewEncoder(w).Encode(http.StatusOK)

	}

	defer myDBpckg.CloseConnectionToDB(db)
}

// Leggi Paziente -
//	- Get Pazient By -
//		- Codice Fiscale:
func getPazientByCodFisc(w http.ResponseWriter, r *http.Request) {
	var codFisc struct {
		CodFisc string `json:"codFisc"`
	}
	err := json.NewDecoder(r.Body).Decode(&codFisc)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	var paziente models.Paziente

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, patientOf " +
		"FROM utente INNER JOIN paziente USING (codFisc) WHERE " +
		"codFisc = " + "'" + codFisc.CodFisc + "' LIMIT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	var tempList string
	// Scan() non può essere eseguita se prima non viene invocata Next()
	for query.Next() {
		// chiaramente nella query non ho inserito la password; il campo nella struct sarà una stringa vuota
		query.Scan(&paziente.Utente.CodFisc, &paziente.Utente.Nome, &paziente.Utente.Cognome,
			&paziente.Utente.Email, &paziente.Utente.Citta,
			&paziente.Utente.Cellulare, &paziente.Utente.Genere, &tempList)
	}

	// TODO: testare PatientOf[], perché sicuramente il bastardo lo restituisce come una stringa
	// 		te quindi forse conviene creare una variabile stringa temporanea e poi pusharla nell'array
	paziente.PatientOf = utils.GenerateArray(&tempList)
	defer query.Close()

	json.NewEncoder(w).Encode(paziente)
}

//		- Mail:
func getPazientByMail(w http.ResponseWriter, r *http.Request) {
	var email struct {
		Email string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	var paziente models.Paziente

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, patientOf " +
		"FROM utente INNER JOIN paziente USING (codFisc) WHERE " +
		"email = " + "'" + email.Email + "' LIMIT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var tempList string
	// Scan() non può essere eseguita se prima non viene invocata Next()
	for query.Next() {
		// chiaramente nella query non ho inserito la password; il campo nella struct sarà una stringa vuota
		query.Scan(&paziente.Utente.CodFisc, &paziente.Utente.Nome, &paziente.Utente.Cognome,
			&paziente.Utente.Email, &paziente.Utente.Citta,
			&paziente.Utente.Cellulare, &paziente.Utente.Genere, &tempList)
	}

	// TODO: testare PatientOf[], perché sicuramente il bastardo lo restituisce come una stringa
	// 		te quindi forse conviene creare una variabile stringa temporanea e poi pusharla nell'array
	paziente.PatientOf = utils.GenerateArray(&tempList)
	defer query.Close()

	json.NewEncoder(w).Encode(paziente)
}

//		- Numero Telefono:
func getPazientByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	var cellulare struct {
		Cellulare string `json:"cellulare"`
	}
	err := json.NewDecoder(r.Body).Decode(&cellulare)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	var paziente models.Paziente

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, patientOf " +
		"FROM utente INNER JOIN paziente USING (codFisc) WHERE " +
		"cellulare = " + "'" + cellulare.Cellulare + "' LIMIT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var tempList string
	// Scan() non può essere eseguita se prima non viene invocata Next()
	for query.Next() {
		// chiaramente nella query non ho inserito la password; il campo nella struct sarà una stringa vuota
		query.Scan(&paziente.Utente.CodFisc, &paziente.Utente.Nome, &paziente.Utente.Cognome,
			&paziente.Utente.Email, &paziente.Utente.Citta,
			&paziente.Utente.Cellulare, &paziente.Utente.Genere, &tempList)
	}

	// TODO: testare PatientOf[], perché sicuramente il bastardo lo restituisce come una stringa
	// 		te quindi forse conviene creare una variabile stringa temporanea e poi pusharla nell'array
	paziente.PatientOf = utils.GenerateArray(&tempList)
	defer query.Close()

	json.NewEncoder(w).Encode(paziente)
}

// 	- Get All Pazients:
func getAllPatients(w http.ResponseWriter, r *http.Request) {
	allPatients := new([]models.Paziente)
	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)
	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, patientOf FROM utente INNER JOIN paziente USING (codFisc)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var i int = 0
	for query.Next() {
		paziente := new(models.Paziente)
		var tempList string
		query.Scan(&paziente.Utente.CodFisc, &paziente.Utente.Nome, &paziente.Utente.Cognome,
			&paziente.Utente.Email, &paziente.Utente.Citta,
			&paziente.Utente.Cellulare, &paziente.Utente.Genere, &tempList)
		paziente.PatientOf = utils.GenerateArray(&tempList)
		*allPatients = append(*allPatients, *paziente)
		i++
	}

	defer query.Close()

	json.NewEncoder(w).Encode(allPatients)

}

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
