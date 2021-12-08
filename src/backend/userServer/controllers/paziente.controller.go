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
	pazientRouter.HandleFunc("/addpsicologobyemail", addPsicologoByEmail).Methods("PUT")
	pazientRouter.HandleFunc("/removepsicologobyemail", removePsicologoByEmail).Methods("PUT")
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

	decryptCodFisc, _ := utils.DecryptBasic(paziente.Utente.CodFisc)
	decryptEmail, _ := utils.DecryptBasic(paziente.Utente.Email)
	decryptPassword, _ := utils.DecryptBasic(paziente.Utente.Password)

	paziente.Utente.CodFisc = decryptCodFisc[0]
	paziente.Utente.Email = decryptEmail[0]
	paziente.Utente.Password = decryptPassword[0]

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

// ---- Associazione a Psicologi ----

// aggiungi Psicologo
//	- By Email

/* COMMENTO IMPORTANTE:
In realtà questa funzione "fa due cose", e non solo una come dovrebbe.
Nello specifico, esegue sia l'update della lista del paziente (che indica di quanti psicologi è paziente)
sia l'update della lista dei pazienti dello psicologo che è stato aggiunto nella lista del paziente:

paziente (esegue addPsicologo per PsicologoY) -----> pazienteX.patientOf[..., "psicologoY"]
psicologoY -----> psicologoY.pazienti[..., "pazienteX"]

Il motivo consiste nella semplicità di realizzazione: per eseguire entrambe le funzionalità
occorrerebbe creare un sistema di notifiche, ad esempio utilizzando un pattern Observer/Observable,
in cui lo psicologo viene notificato della richiesta di aggiunta di un paziente alla propria lista,
e solo nel caso in cui lo psicologo accetti si eseguano le operazioni sopraelencate.

Il discorso è analogo per la deletePsicologoByEmail (funzione successiva)
*/
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
		// quando ne si fa una GET (vedi funzioni soprastanti), la stringa "patientOf" viene parsata come array grazie alla funzione che ho scritto in utils/arrayGen.go
	}

	// eseguire controllo su stringa (da trasformare quindi in array) per vedere se l'email è già presente:
	if strings.Contains(elencoEmailPsicologiStringa, addInfo.Email) {
		http.Error(w, "Psicologo già presente in lista patientOf", http.StatusMethodNotAllowed)

		// !!!!!! se non si chiude la connessione col database in maniera esplicita, le query di update vengono comunque eseguite !!!!!!
		myDBpckg.CloseConnectionToDB(db)
		return
	}

	elencoEmailPsicologiStringa = addInfo.Email + "," + elencoEmailPsicologiStringa
	defer queryListaPsicologi.Close()

	// query per aggiornare il paziente
	queryUpdatePaziente, err := db.Query("UPDATE paziente SET patientOf=" + "'" + elencoEmailPsicologiStringa /* strings.Join(elencoEmailPsicologi, ",") */ + "' WHERE " +
		"codFisc=" + "'" + addInfo.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryUpdatePaziente.Close()

	// parte per aggiungere l'utente alla lista dei pazienti dello psicologo
	var listaPazienti string
	var codFiscPsicologo string // Lo uso per la update dello psicologo, dopo che ho ottenuto la coppia codFisc + pazienti dalla query successiva

	queryGetPsicologoByEmail, err := db.Query("SELECT codFisc, pazienti FROM utente INNER JOIN psicologo USING (codFisc) WHERE utente.email='" + addInfo.Email + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryGetPsicologoByEmail.Close()

	for queryGetPsicologoByEmail.Next() {
		queryGetPsicologoByEmail.Scan(&codFiscPsicologo, &listaPazienti)
	}

	// controllo se lo psicologo ha già il codice fiscale del paziente inserito nella propria lista:
	if strings.Contains(listaPazienti, addInfo.CodFisc) {
		http.Error(w, "Paziente già presente in lista pazienti", http.StatusMethodNotAllowed)
		myDBpckg.CloseConnectionToDB(db)
		return
	}

	listaPazienti = addInfo.CodFisc + "," + listaPazienti

	queryUpdatePsicologo, err := db.Query("UPDATE psicologo SET pazienti='" + listaPazienti + "' WHERE psicologo.codFisc='" + codFiscPsicologo + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryUpdatePsicologo.Close()
}

// rimuovi Psicologo
func removePsicologoByEmail(w http.ResponseWriter, r *http.Request) {
	var removeInfo struct {
		CodFisc string `json:"codFisc"` // codice fiscale del paziente
		Email   string `json:"email`    // email dello psicologo
	}
	err := json.NewDecoder(r.Body).Decode(&removeInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()
	defer myDBpckg.CloseConnectionToDB(db)

	// ottieni la lista patientOf del paziente via codFisc
	queryGetPatientOf, err := db.Query("SELECT patientOf FROM paziente WHERE codFisc='" + removeInfo.CodFisc + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryGetPatientOf.Close()
	var patientOf string
	for queryGetPatientOf.Next() {
		queryGetPatientOf.Scan(&patientOf)
	}

	// controlla se l'email dello psicologo è presente
	// se lo psicologo non è presente nella lista patientOf, chiudiamo connessione col db e returniamo
	if !strings.Contains(patientOf, removeInfo.Email) {
		http.Error(w, "Psicologo non presente in lista", http.StatusNotFound)
		myDBpckg.CloseConnectionToDB(db)
		return
	}

	// altrimenti rimuovi email dello psicologo
	patientOf = strings.ReplaceAll(patientOf, removeInfo.Email+",", "")

	// UPDATE
	queryUpdatePatientOf, err := db.Query("UPDATE paziente SET patientOf='" + patientOf + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer queryUpdatePatientOf.Close()

	// ottieni la lista pazienti dello psicologo (SELECT via email)
	var listaPazienti string
	var codFiscPsicologo string

	queryGetPsicologoByEmail, err := db.Query("SELECT codFisc, pazienti FROM utente INNER JOIN psicologo USING (codFisc) WHERE utente.email='" + removeInfo.Email + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryGetPsicologoByEmail.Close()

	for queryGetPsicologoByEmail.Next() {
		queryGetPsicologoByEmail.Scan(&codFiscPsicologo, &listaPazienti)
	}

	// controlla se il codice fiscale del paziente è presente nella lista
	if !strings.Contains(listaPazienti, removeInfo.CodFisc) {
		http.Error(w, "Paziente non presente in lista pazienti", http.StatusNotFound)
		myDBpckg.CloseConnectionToDB(db)
		return
	}

	// rimuovi il codice fiscale del paziente
	listaPazienti = strings.ReplaceAll(listaPazienti, removeInfo.CodFisc+",", "")

	// UPDATE
	queryUpdatePsicologo, err := db.Query("UPDATE psicologo SET pazienti='" + listaPazienti + "' WHERE codFisc='" + codFiscPsicologo + "';")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer queryUpdatePsicologo.Close()
}
