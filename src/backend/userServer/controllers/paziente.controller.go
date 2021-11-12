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
	pazientRouter.HandleFunc("/getbyemail", getPazientByEmail).Methods("GET")
	pazientRouter.HandleFunc("/getbyphonenumber", getPazientByPhoneNumber).Methods("GET")
	pazientRouter.HandleFunc("/getallpatients", getAllPatients).Methods("GET")
	pazientRouter.HandleFunc("/updatephonenumber", updateParientPhoneNumber).Methods("PUT")
	pazientRouter.HandleFunc("/updateemail", updateParientEmail).Methods("PUT")
	pazientRouter.HandleFunc("/deletepatientbyemail", deletePatientByEmail).Methods("DELETE")

	// gestione psicologo:
	pazientRouter.HandleFunc("/addpsicologbyemail", addPsicologoByEmail).Methods("PUT")
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

	paziente := new(models.Paziente)
	err := json.NewDecoder(r.Body).Decode(paziente)

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

	errCheck := db.QueryRow("SELECT utente.codFisc FROM utente INNER JOIN paziente ON utente.codFisc = paziente.codFisc " +
		"WHERE paziente.codFisc = '" + paziente.Utente.CodFisc + "';").Scan(&isPresent)

	if errCheck != nil && errCheck != sql.ErrNoRows {

		http.Error(w, "User already exists" /* errCheck.Error() */, http.StatusForbidden)
		return

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

	paziente := new(models.Paziente)

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, patientOf " +
		"FROM utente INNER JOIN paziente USING (codFisc) WHERE " +
		"codFisc = " + "'" + codFisc.CodFisc + "' LIMIT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	defer query.Close()

	var tempList string
	// Scan() non può essere eseguita se prima non viene invocata Next()
	for query.Next() {
		// chiaramente nella query non ho inserito la password; il campo nella struct sarà una stringa vuota
		query.Scan(&paziente.Utente.CodFisc, &paziente.Utente.Nome, &paziente.Utente.Cognome,
			&paziente.Utente.Email, &paziente.Utente.Citta,
			&paziente.Utente.Cellulare, &paziente.Utente.Genere, &tempList)

	}

	if paziente.Utente.CodFisc == "" {
		http.Error(w, "Paziente non trovato", http.StatusNotFound)
		return
	}

	// TODO: testare PatientOf[], perché sicuramente il bastardo lo restituisce come una stringa
	// 		te quindi forse conviene creare una variabile stringa temporanea e poi pusharla nell'array
	paziente.PatientOf = utils.GenerateArray(&tempList)

	json.NewEncoder(w).Encode(paziente)
}

//		- Mail:
func getPazientByEmail(w http.ResponseWriter, r *http.Request) {
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

	paziente := new(models.Paziente)

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

	if paziente.Utente.CodFisc == "" {
		http.Error(w, "Paziente non trovato", http.StatusNotFound)
		return
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

	paziente := new(models.Paziente)

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, patientOf " +
		"FROM utente INNER JOIN paziente USING (codFisc) WHERE " +
		"cellulare = " + "'" + cellulare.Cellulare + "' LIMIT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer query.Close()

	var tempList string
	// Scan() non può essere eseguita se prima non viene invocata Next()
	for query.Next() {
		// chiaramente nella query non ho inserito la password; il campo nella struct sarà una stringa vuota
		query.Scan(&paziente.Utente.CodFisc, &paziente.Utente.Nome, &paziente.Utente.Cognome,
			&paziente.Utente.Email, &paziente.Utente.Citta,
			&paziente.Utente.Cellulare, &paziente.Utente.Genere, &tempList)
	}

	if paziente.Utente.CodFisc == "" {
		http.Error(w, "Paziente non trovato", http.StatusNotFound)
		return
	}

	// TODO: testare PatientOf[], perché sicuramente il bastardo lo restituisce come una stringa
	// 		te quindi forse conviene creare una variabile stringa temporanea e poi pusharla nell'array
	paziente.PatientOf = utils.GenerateArray(&tempList)

	json.NewEncoder(w).Encode(paziente)
}

// 	- Get All Pazients:
func getAllPatients(w http.ResponseWriter, r *http.Request) {
	allPatients := new([]models.Paziente)
	paziente := new(models.Paziente)
	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, patientOf FROM utente INNER JOIN paziente USING (codFisc);")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer query.Close()

	for query.Next() {
		var tempList string
		query.Scan(&paziente.Utente.CodFisc, &paziente.Utente.Nome, &paziente.Utente.Cognome,
			&paziente.Utente.Email, &paziente.Utente.Citta,
			&paziente.Utente.Cellulare, &paziente.Utente.Genere, &tempList)
		paziente.PatientOf = utils.GenerateArray(&tempList)
		*allPatients = append(*allPatients, *paziente)
	}

	// controlla la dimensione dell'array; se è vuoto -> nessun paziente trovato
	if len(*allPatients) == 0 {
		http.Error(w, "Nessun paziente trovato", http.StatusNotFound)
		return

	}

	json.NewEncoder(w).Encode(allPatients)
}

// Update Paziente
//	- Update Email:
func updateParientEmail(w http.ResponseWriter, r *http.Request) {
	var updateStruct struct {
		CodFisc    string `json:"codFisc"`
		NuovaEmail string `json:"nuovaEmail`
	}
	err := json.NewDecoder(r.Body).Decode(&updateStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := myDBpckg.ConnectToDB()
	defer db.Close()

	/**EFFETTUARE UN QUERY CHECK SULL'ESISTENZA DELL'UTENTE VIA CODFISC
	 * (anche se in realtà non è necessario, dato che l'update viene eseguita sul nulla e non genera errore)
	 */
	query, err := db.Query("UPDATE utente SET email=" + "'" + updateStruct.NuovaEmail + "' WHERE " +
		"codFisc=" + "'" + updateStruct.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer query.Close()
}

//	- Update num Tel:
func updateParientPhoneNumber(w http.ResponseWriter, r *http.Request) {
	var updateStruct struct {
		CodFisc        string `json:"codFisc"`
		NuovoCellulare string `json:"nuovoCellulare`
	}
	err := json.NewDecoder(r.Body).Decode(&updateStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := myDBpckg.ConnectToDB()
	defer db.Close()

	/**EFFETTUARE UN QUERY CHECK SULL'ESISTENZA DELL'UTENTE VIA CODFISC
	 * (anche se in realtà non è necessario, dato che l'update viene eseguita sul nulla e non genera errore)
	 */
	query, err := db.Query("UPDATE utente SET cellulare=" + "'" + updateStruct.NuovoCellulare + "' WHERE " +
		"codFisc=" + "'" + updateStruct.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer query.Close()
}

//	- Update Genere: -- NEXT USE CASE

// Rimuovi Paziente:
//	- per email
func deletePatientByEmail(w http.ResponseWriter, r *http.Request) {
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
	query, err := db.Query("DELETE utente, paziente FROM utente INNER JOIN paziente ON utente.codFisc=paziente.codFisc WHERE " +
		"utente.email=" + "'" + email.Email + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer query.Close()
}

// ---- CRUD Associazione a Psicologi ----

// aggiungi Psicologo
//	- By Email
func addPsicologoByEmail(w http.ResponseWriter, r *http.Request) {
	var addInfo struct {
		CodFisc string `json:"codFisc"` // codice fiscale del paziente
		Email   string `json:"email`    // email dello psicologo
	}
	err := json.NewDecoder(r.Body).Decode(&addInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	// var elencoEmailPsicologi []string -> noi la otteniamo dalla query come stringa
	var elencoEmailPsicologiStringa string
	queryListaPsicologi, err := db.Query("SELECT patientOf FROM paziente WHERE codFisc='" + addInfo.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	for queryListaPsicologi.Next() {
		queryListaPsicologi.Scan(&elencoEmailPsicologiStringa)
		// quando ne si fa una GET, la stringa "patientOf" viene parsata come array grazie alla funzione che ho scritto in utils/arrayGen.go
	}

	// eseguire controllo su stringa (da trasformare quindi in array) per vedere se l'email è già presente?

	elencoEmailPsicologiStringa = elencoEmailPsicologiStringa + "," + addInfo.Email
	defer queryListaPsicologi.Close()

	// query per aggiornare il paziente
	queryUpdatePaziente, err := db.Query("UPDATE paziente SET patientOf=" + "'" + elencoEmailPsicologiStringa /* strings.Join(elencoEmailPsicologi, ",") */ + "' WHERE " +
		"codFisc=" + "'" + addInfo.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	defer queryUpdatePaziente.Close()
}

// rimuovi Psicologo
