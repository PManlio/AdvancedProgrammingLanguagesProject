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

func PsicologoHandler(psicologoRouter *mux.Router) {
	// ping-pong di /psicologo
	psicologoRouter.HandleFunc("/pong", pong).Methods("GET")

	// CRUD Psicologo
	psicologoRouter.HandleFunc("/create", CreatePsicologo).Methods("POST")
	psicologoRouter.HandleFunc("/getbycodfisc", getPsicologoByCodFisc).Methods("GET")
	psicologoRouter.HandleFunc("/getbyemail", getPsicologoByEmail).Methods("GET")
	psicologoRouter.HandleFunc("/getbyphonenumber", getPsicologoByPhoneNumber).Methods("GET")
	psicologoRouter.HandleFunc("/getallpsicologi", getAllPsicologi).Methods("GET")
	psicologoRouter.HandleFunc("/getallpsicologibycity", getPsicologiByCity).Methods("GET")
	psicologoRouter.HandleFunc("/updatephonenumber", updatePsicologoPhoneNumber).Methods("PUT")
	psicologoRouter.HandleFunc("/updateemail", updatePsicologoEmail).Methods("PUT")
	psicologoRouter.HandleFunc("/deletepsicologobyemail", deletePsicologoByEmail).Methods("DELETE")
}

// al solito, semplice ping per testare routing
/* qui è chiamata pong perché all'interno del package
 * in psicologo.controller.go è già definita la funzione
 * ping e quindi qui la chiamiamo pong per evitare warning
**/
func pong(w http.ResponseWriter, r *http.Request) {
	pong := struct {
		Pong string
	}{
		Pong: "ping",
	}
	json.NewEncoder(w).Encode(pong)
}

// ---- CRUD Psicologo ----

// Crea Psicologo:
func CreatePsicologo(w http.ResponseWriter, r *http.Request) {

	// var psicologo models.Psicologo
	psicologo := new(models.Psicologo)
	err := json.NewDecoder(r.Body).Decode(&psicologo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Println(psicologo)
	psicologo.Utente.Password = utils.Encrypt(psicologo.Utente.Password)
	// fmt.Println(psicologo)

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	// CONTROLLO SU TABELLA PER VERIFICARE ESISTENZA
	var isPresent bool

	// forse questo err check andrebbe fatto solo con codFisc su UTENTE, non suna innerJoin
	errCheck := db.QueryRow("SELECT utente.codFisc FROM utente INNER JOIN psicologo ON utente.codFisc = psicologo.codFisc " +
		"WHERE psicologo.codFisc = '" + psicologo.Utente.CodFisc + "';").Scan(&isPresent)

	if errCheck != nil && errCheck != sql.ErrNoRows {

		http.Error(w, "User already exists" /* errCheck.Error() */, http.StatusForbidden)
		return

	} else if !isPresent {

		insert1, err1 := db.Query("INSERT INTO utente(codFisc, nome, cognome, email, password, citta, cellulare, genere) " +
			"VALUES(" + "'" + psicologo.Utente.CodFisc + "'" + ", " + "'" + psicologo.Utente.Nome + "'" +
			", " + "'" + psicologo.Utente.Cognome + "'" + ", " + "'" + psicologo.Utente.Email + "'" + ", " +
			"'" + psicologo.Utente.Password + "'" + ", " + "'" + psicologo.Utente.Citta + "'" + ", " +
			"'" + psicologo.Utente.Cellulare + "'" + ", " + "'" + psicologo.Utente.Genere + "');")
		if err1 != nil {
			fmt.Println("--------\n\nun errore di query è avvenuto:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer insert1.Close()

		insert2, err2 := db.Query("INSERT INTO psicologo(codFisc, pazienti) VALUES(" + "'" + psicologo.Utente.CodFisc + "'" + ", " +
			"'" + strings.Join(psicologo.Pazienti, ",") + "');")
		if err2 != nil {
			fmt.Println("--------\n\nun errore di query è avvenuto:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer insert2.Close()

		json.NewEncoder(w).Encode(http.StatusOK)
	}
}

// Leggi psicologo -
//	- Get Pazient By -
//		- Codice Fiscale:
func getPsicologoByCodFisc(w http.ResponseWriter, r *http.Request) {
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

	var psicologo models.Psicologo

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, Pazienti " +
		"FROM utente INNER JOIN psicologo USING (codFisc) WHERE " +
		"codFisc = " + "'" + codFisc.CodFisc + "' LIMIT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	var tempList string
	// Scan() non può essere eseguita se prima non viene invocata Next()
	for query.Next() {
		// chiaramente nella query non ho inserito la password; il campo nella struct sarà una stringa vuota
		query.Scan(&psicologo.Utente.CodFisc, &psicologo.Utente.Nome, &psicologo.Utente.Cognome,
			&psicologo.Utente.Email, &psicologo.Utente.Citta,
			&psicologo.Utente.Cellulare, &psicologo.Utente.Genere, &tempList)

	}

	if psicologo.Utente.CodFisc == "" {
		http.Error(w, "psicologo non trovato", http.StatusNotFound)
		return
	}

	// TODO: testare Pazienti[], perché sicuramente il bastardo lo restituisce come una stringa
	// 		te quindi forse conviene creare una variabile stringa temporanea e poi pusharla nell'array
	psicologo.Pazienti = utils.GenerateArray(&tempList)
	defer query.Close()

	json.NewEncoder(w).Encode(psicologo)
}

//		- Mail:
func getPsicologoByEmail(w http.ResponseWriter, r *http.Request) {
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

	var psicologo models.Psicologo

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, Pazienti " +
		"FROM utente INNER JOIN psicologo USING (codFisc) WHERE " +
		"email = " + "'" + email.Email + "' LIMIT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var tempList string
	// Scan() non può essere eseguita se prima non viene invocata Next()
	for query.Next() {
		// chiaramente nella query non ho inserito la password; il campo nella struct sarà una stringa vuota
		query.Scan(&psicologo.Utente.CodFisc, &psicologo.Utente.Nome, &psicologo.Utente.Cognome,
			&psicologo.Utente.Email, &psicologo.Utente.Citta,
			&psicologo.Utente.Cellulare, &psicologo.Utente.Genere, &tempList)
	}

	if psicologo.Utente.CodFisc == "" {
		http.Error(w, "psicologo non trovato", http.StatusNotFound)
		return
	}

	// TODO: testare Pazienti[], perché sicuramente il bastardo lo restituisce come una stringa
	// 		te quindi forse conviene creare una variabile stringa temporanea e poi pusharla nell'array
	psicologo.Pazienti = utils.GenerateArray(&tempList)
	defer query.Close()

	json.NewEncoder(w).Encode(psicologo)
}

//		- Numero Telefono:
func getPsicologoByPhoneNumber(w http.ResponseWriter, r *http.Request) {
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

	var psicologo models.Psicologo

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, Pazienti " +
		"FROM utente INNER JOIN psicologo USING (codFisc) WHERE " +
		"cellulare = " + "'" + cellulare.Cellulare + "' LIMIT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var tempList string
	// Scan() non può essere eseguita se prima non viene invocata Next()
	for query.Next() {
		// chiaramente nella query non ho inserito la password; il campo nella struct sarà una stringa vuota
		query.Scan(&psicologo.Utente.CodFisc, &psicologo.Utente.Nome, &psicologo.Utente.Cognome,
			&psicologo.Utente.Email, &psicologo.Utente.Citta,
			&psicologo.Utente.Cellulare, &psicologo.Utente.Genere, &tempList)
	}

	if psicologo.Utente.CodFisc == "" {
		http.Error(w, "psicologo non trovato", http.StatusNotFound)
		return
	}

	// TODO: testare Pazienti[], perché sicuramente il bastardo lo restituisce come una stringa
	// 		te quindi forse conviene creare una variabile stringa temporanea e poi pusharla nell'array
	psicologo.Pazienti = utils.GenerateArray(&tempList)
	defer query.Close()

	json.NewEncoder(w).Encode(psicologo)
}

type PsicoInfo struct {
	Nome, Cognome, Email, Citta, Cellulare, Genere string
}

func getPsicologiByCity(w http.ResponseWriter, r *http.Request) {
	var citta struct {
		Citta string `json:"citta"`
	}
	err := json.NewDecoder(r.Body).Decode(&citta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)
	listaInfoPsicologi := new([]PsicoInfo)
	infoPsicologo := new(PsicoInfo)

	query, err := db.Query("SELECT nome, cognome, email, citta, cellulare, genere FROM utente INNER JOIN psicologo USING (codFisc) WHERE utente.citta=" +
		"'" + citta.Citta + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer query.Close()

	for query.Next() {
		query.Scan(&infoPsicologo.Nome, &infoPsicologo.Cognome, &infoPsicologo.Email,
			&infoPsicologo.Citta, &infoPsicologo.Cellulare, &infoPsicologo.Genere)
		*listaInfoPsicologi = append(*listaInfoPsicologi, *infoPsicologo)
	}

	if len(*listaInfoPsicologi) == 0 {
		http.Error(w, "Nessun psicologo trovato", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(listaInfoPsicologi)

}

// 	- Get All Psicologi:
func getAllPsicologi(w http.ResponseWriter, r *http.Request) {
	allPsicologi := new([]models.Psicologo)
	psicologo := new(models.Psicologo)
	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	query, err := db.Query("SELECT codFisc, nome, cognome, email, citta, cellulare, genere, Pazienti FROM utente INNER JOIN psicologo USING (codFisc);")

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer query.Close()

	for query.Next() {
		var tempList string
		query.Scan(&psicologo.Utente.CodFisc, &psicologo.Utente.Nome, &psicologo.Utente.Cognome,
			&psicologo.Utente.Email, &psicologo.Utente.Citta,
			&psicologo.Utente.Cellulare, &psicologo.Utente.Genere, &tempList)
		psicologo.Pazienti = utils.GenerateArray(&tempList)
		*allPsicologi = append(*allPsicologi, *psicologo)
	}

	// controlla la dimensione dell'array; se è vuoto -> nessun psicologo trovato
	if len(*allPsicologi) == 0 {
		http.Error(w, "Nessun psicologo trovato", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(allPsicologi)

}

// Update psicologo
//	- Update Email:
func updatePsicologoEmail(w http.ResponseWriter, r *http.Request) {
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
func updatePsicologoPhoneNumber(w http.ResponseWriter, r *http.Request) {
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

// Rimuovi psicologo:
//	- per email
func deletePsicologoByEmail(w http.ResponseWriter, r *http.Request) {
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
	query, err := db.Query("DELETE utente, psicologo FROM utente INNER JOIN psicologo ON utente.codFisc=psicologo.codFisc WHERE " +
		"utente.email=" + "'" + email.Email + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer query.Close()
}

// aggiungi Paziente

// leggi Paziente -
//	- Get Paziente By CodFisc
//	- Get All Patients dello psicologo che ne effettua la richiesta

// rimuovi Paziente
