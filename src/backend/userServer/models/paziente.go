package models

type Paziente struct {
	Utente    Utente   `json:"utente"`
	PatientOf []string `json:"patientOf"`
}
