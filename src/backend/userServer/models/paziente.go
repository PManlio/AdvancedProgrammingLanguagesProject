package models

type Paziente struct {
	Utente    Utente   `json:"utente"`
	PatientOf []string `json:"patientOf"`
}

type PazienteInfo struct {
	CodFisc, Nome, Cognome, Email, Cellulare, Genere string
}
