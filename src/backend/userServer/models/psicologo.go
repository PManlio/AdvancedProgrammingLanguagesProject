package models

type Psicologo struct {
	Utente   Utente   `json:"utente"`
	Pazienti []string `json:"pazienti"`
}
