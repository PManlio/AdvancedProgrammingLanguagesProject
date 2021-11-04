package models

type Psicologo struct {
	Utente   Utente     `json:"utente"`
	Pazienti []Paziente `json:"pazienti"`
}
