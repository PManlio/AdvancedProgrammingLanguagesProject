package manager

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	myDBpckg "../db"
	"../utils"

	"github.com/gorilla/mux"
)

func LoginHandler(loginRouter *mux.Router) {
	loginRouter.HandleFunc("", Login).Methods("POST")
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "" {
		http.Error(w, "No Basic Authorization", http.StatusUnauthorized)
		return
	}
	reqToken := strings.Split(r.Header.Get("Authorization"), "Basic ")[1]
	if reqToken == "" {
		http.Error(w, "no bearer in header", http.StatusBadRequest)
		return
	}

	userInfo, err := utils.DecryptBasic(reqToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if (userInfo[0] == "") || (userInfo[1] == "") {
		http.Error(w, "No Email or Password present", http.StatusUnauthorized)
		return
	}

	email := userInfo[0]
	password := userInfo[1]

	db := myDBpckg.ConnectToDB()

	var hashpassword string = utils.Encrypt(password)

	queryFindUser, err := db.Query("SELECT codFisc FROM utente WHERE email='" + email + "' AND password='" + hashpassword + "';")
	if err != nil {
		http.Error(w, "utente non trovato", http.StatusUnauthorized)
		return
	}

	if !queryFindUser.Next() {
		http.Error(w, "utente non trovato", http.StatusUnauthorized)
		return
	}

	// per ora lo tengo come struct
	utente := new(utils.Jwt)
	for queryFindUser.Next() {
		queryFindUser.Scan(&utente.CodFisc)
		utente.Date = time.Now()
	}

	token, err := utente.GenerateToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer myDBpckg.CloseConnectionToDB(db)

	var returnJwt struct {
		Rjwt string `json:"token"`
	}
	returnJwt.Rjwt = token

	json.NewEncoder(w).Encode(returnJwt)
}
