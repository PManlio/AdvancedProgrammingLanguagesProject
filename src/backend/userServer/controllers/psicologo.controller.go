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

	// gestione pazienti
	psicologoRouter.HandleFunc("/addpazientebyemail", addPazienteByEmail).Methods("PUT")
	psicologoRouter.HandleFunc("/getpazienti", getPazientiPsicologo).Methods("GET")
	psicologoRouter.HandleFunc("/removepaziente", removePazienteByCodFisc).Methods("PUT")
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

	psicologo := new(models.Psicologo)

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

	psicologo := new(models.Psicologo)

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

	psicologo := new(models.Psicologo)

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

// prende tutti gli psicologi con una determinata città
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
	listaInfoPsicologi := new([]models.Utente)
	infoPsicologo := new(models.Utente)

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

// aggiungi Paziente (ne salvo il codice fiscale)
func addPazienteByEmail(w http.ResponseWriter, r *http.Request) {
	var addInfo struct {
		Email   string `json:"email"`   // email Paziente
		CodFisc string `json:"codFisc"` // codice fiscale Psicologo
	}
	err := json.NewDecoder(r.Body).Decode(&addInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	// query per aggiungere il codice fiscale del paziente nell'elenco pazienti dello psicologo
	queryAddPaziente, err := db.Query("SELECT pazienti FROM psicologo WHERE codFisc='" + addInfo.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryAddPaziente.Close()

	var listaPazienti, nuovoPaziente string // nuovoPaziente contiene il codice fiscale del paziente da aggiungere
	for queryAddPaziente.Next() {
		queryAddPaziente.Scan(&listaPazienti)
	}

	// prendo il codice fiscale del paziente tramite la sua email
	querySelectPaziente, err := db.Query("SELECT codFisc FROM utente WHERE " +
		"email='" + addInfo.Email + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer querySelectPaziente.Close()
	for querySelectPaziente.Next() {
		querySelectPaziente.Scan(&nuovoPaziente)
	}

	// controllo se il paziente è già presente in elenco
	if strings.Contains(listaPazienti, nuovoPaziente) {
		http.Error(w, "paziente già presente in lista pazienti", http.StatusMethodNotAllowed)
		myDBpckg.CloseConnectionToDB(db)
		return
	}

	// UPDATE
	listaPazienti = nuovoPaziente + "," + listaPazienti
	queryUpdatePazienti, err := db.Query("UPDATE psicologo SET pazienti='" + listaPazienti + "' WHERE codFisc='" + addInfo.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryUpdatePazienti.Close()

	// query per aggiornare l'elenco patientOf del paziente aggiunto
	var listaPatientOf string
	queryGetPatientOf, err := db.Query("SELECT patientOf FROM paziente WHERE codFisc='" + nuovoPaziente + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	defer queryGetPatientOf.Close()
	for queryGetPatientOf.Next() {
		queryGetPatientOf.Scan(&listaPatientOf)
	}

	var emailPsicologo string
	queryGetEmailPsicologo, err := db.Query("SELECT email FROM utente WHERE codFisc='" + addInfo.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	defer queryGetEmailPsicologo.Close()
	for queryGetEmailPsicologo.Next() {
		queryGetEmailPsicologo.Scan(&emailPsicologo)
	}

	listaPatientOf = listaPatientOf + "," + emailPsicologo
	queryUpdatePaziente, err := db.Query("UPDATE paziente SET patientOf='" + listaPatientOf + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	defer queryUpdatePaziente.Close()
}

// --------- GESTIONE LISTA PAZIENTI PSICOLOGO ---------

// get Paziente -
//	- Get pazienti dello psicologo
func getPazientiPsicologo(w http.ResponseWriter, r *http.Request) {
	var psicologo struct {
		CodFisc string `json:"codFisc"`
	}
	err := json.NewDecoder(r.Body).Decode(&psicologo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	queryGetLista, err := db.Query("SELECT pazienti FROM psicologo WHERE codFisc='" + psicologo.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryGetLista.Close()

	var listaString string
	for queryGetLista.Next() {
		queryGetLista.Scan(&listaString)
	}

	var listaArray []string = utils.GenerateArray(&listaString)

	if len(listaArray) == 0 {
		http.Error(w, "Nessun paziente in lista", http.StatusNotFound)
		myDBpckg.CloseConnectionToDB(db)
		return
	}

	pazienti := new([]models.Utente)
	paziente := new(models.Utente)

	for _, p := range listaArray {
		queryPaziente, err := db.Query("SELECT codFisc, nome, cognome, email, cellulare, genere FROM utente WHERE codFisc='" + p + "';")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer queryPaziente.Close()

		for queryPaziente.Next() {
			queryPaziente.Scan(&paziente.CodFisc, &paziente.Nome, &paziente.Cognome, &paziente.Email,
				&paziente.Cellulare, &paziente.Genere)
		}
		*pazienti = append(*pazienti, *paziente)
	}

	json.NewEncoder(w).Encode(pazienti)
}

//	- Get Paziente By codFisc già c'è in paziente.controller

// rimuovi Paziente
func removePazienteByCodFisc(w http.ResponseWriter, r *http.Request) {
	var info struct {
		CodFiscPsicologo string `json:"codFiscPsicologo"`
		CodFiscPaziente  string `json:"codFiscPaziente"`
	}
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	queryGetPazienti, err := db.Query("SELECT pazienti FROM psicologo WHERE codFisc='" + info.CodFiscPsicologo + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryGetPazienti.Close()

	var listaPazienti string
	for queryGetPazienti.Next() {
		queryGetPazienti.Scan(&listaPazienti)
	}
	listaPazienti = strings.ReplaceAll(listaPazienti, info.CodFiscPaziente+",", "")

	queryUpdatePazienti, err := db.Query("UPDATE psicologo SET pazienti='" + listaPazienti + "' WHERE codFisc='" + info.CodFiscPsicologo + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	defer queryUpdatePazienti.Close()

	queryGetEmailPsicologo, err := db.Query("SELECT email FROM utente WHERE codFisc='" + info.CodFiscPsicologo + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryGetEmailPsicologo.Close()

	var emailPsicologo string
	for queryGetEmailPsicologo.Next() {
		queryGetEmailPsicologo.Scan(&emailPsicologo)
	}

	queryGetPaziente, err := db.Query("SELECT patientOf FROM paziente WHERE codFisc='" + info.CodFiscPaziente + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryGetPaziente.Close()

	var listaPatientOf string
	for queryGetPaziente.Next() {
		queryGetPaziente.Scan(&listaPatientOf)
	}
	listaPatientOf = strings.ReplaceAll(listaPatientOf, emailPsicologo+",", "")

	queryUpdatePatientOf, err := db.Query("UPDATE paziente SET patientOf='" + listaPatientOf + "' WHERE codFisc='" + info.CodFiscPaziente + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	defer queryUpdatePatientOf.Close()
}
