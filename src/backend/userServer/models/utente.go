package models

type Utente struct {
	CodFisc   string `json: "codFisc"`
	Nome      string `json: "nome"`
	Cognome   string `json: "cognome"`
	Email     string `json: "emai"`
	Password  string `json: "password"`
	Citta     string `json: "citta"`
	Cellulare string `json: "cellulare"`
	Genere    string `json: "genere"`
}
