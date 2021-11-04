package models

type Paziente struct {
	Utente    Utente
	PatientOf []string
}
