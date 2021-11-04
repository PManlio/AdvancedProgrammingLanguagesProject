package controllers

import (
	"encoding/json"
	"net/http"

	// sql "../db"

	"github.com/gorilla/mux"
)

func PazientHandler(pazientRouter *mux.Router) {
	pazientRouter.HandleFunc("/ping", ping)
}

// ---- CRUD Paziente ----

// simple Ping
func ping(w http.ResponseWriter, r *http.Request) {
	ping := struct {
		Ping string
	}{
		Ping: "pong",
	}
	json.NewEncoder(w).Encode(ping)
}

// Crea Paziente:
//func (p PazientController) CreatePazient()

// Leggi Paziente -
//	- Get Pazient By -
//		- Codice Fiscale:
//		- Mail:
//		- Numero Telefono:

// 	- Get All Pazients:

// Update Paziente
//	- Update Email:

//	- Update num Tel:

//	- Update Genere:

// Rimuovi Paziente:

// ---- CRUD Associazione a Psicologi ----

// aggiungi Psicologo

// leggi Psicologo -
//	- Get Psicologo By CodFisc
//	- Get All Psicologysts

// rimuovi Psicologo
