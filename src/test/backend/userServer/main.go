package main

import (
	"fmt"

	"./models"
)

func main() {

	fmt.Println("esisto")

	paziente := &models.Paziente{
		Utente: &models.Utente{
			CodFisc:   "12345",
			Nome:      "aeiou",
			Cognome:   "yeyr",
			Email:     "string",
			Citta:     "string",
			Cellulare: "string",
			Genere:    "string",
		},
		PatientOf: []string{"Prova"},
	}

	paziente2 := &models.Paziente{
		/*Utente: models.Utente{
			CodFisc:   "67890",
			Nome:      "uoiea",
			Cognome:   "ieij",
			Email:     "string",
			Citta:     "string",
			Cellulare: "string",
			Genere:    "string",
		},
		PatientOf: []string{"Prova"},*/
	}

	psicologo := &models.Psicologo{
		Utente: &models.Utente{
			CodFisc:   "12345",
			Nome:      "aeiou",
			Cognome:   "yeyr",
			Email:     "string",
			Citta:     "string",
			Cellulare: "string",
			Genere:    "string",
		},
		Pazienti: []models.Paziente{*paziente},
	}

	utente := &models.Utente{
		CodFisc:   "23232",
		Nome:      "aqwert",
		Cognome:   "shock",
		Email:     "string",
		Citta:     "string",
		Cellulare: "string",
		Genere:    "string",
	}
	paziente3 := &models.Paziente{
		Utente: utente,
	}

	if &utente == &paziente3.Utente {
		fmt.Println("sì sono uguali")
	}

	fmt.Println(paziente, psicologo, &paziente2, &utente, &paziente3.Utente, &paziente3)

	a := &struct{ k string }{k: "sex"}

	var b = a
	b.k = "on the"
	fmt.Println(a.k, b.k)

	var c = a
	c.k = "beach"
	fmt.Println(a.k, b.k, c.k)

	fmt.Println(&a, &b, &c)

}
