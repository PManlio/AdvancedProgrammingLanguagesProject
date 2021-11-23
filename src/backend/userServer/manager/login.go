package manager

import (
	"encoding/json"
	"net/http"

	myDBpckg "../db"
	"../models"
	"../utils"

	"github.com/gorilla/mux"
)

func LoginHandler(loginRouter *mux.Router) {
	loginRouter.HandleFunc("/login", Login)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// da riscrivere con r.Header["Bearer"] D:
	var loginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := myDBpckg.ConnectToDB()

	var hashpassword string = utils.Encrypt(loginInfo.Password)

	queryFindUser, err := db.Query("SELECT codFisc, nome, cognome, email FROM utente WHERE email='" + loginInfo.Email + "' AND password='" + hashpassword + "';")
	if err != nil {
		http.Error(w, "utente non trovato", http.StatusNotFound)
		return
	}
	utente := new(models.Utente)
	for queryFindUser.Next() {
		queryFindUser.Scan(&utente.CodFisc, &utente.Nome, &utente.Cognome, &utente.Email)
	}

	defer myDBpckg.CloseConnectionToDB(db)
}
