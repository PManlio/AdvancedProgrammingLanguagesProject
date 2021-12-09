package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	myDBpckg "../db"
	"../models"
	"../utils"
)

func checkRequest(w http.ResponseWriter, r *http.Request) (bool, string, error) {
	token := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	if token == "" {
		// http.Error(w, "Missing Token", http.StatusForbidden)
		return false, "", fmt.Errorf("Token is missing.")
	}

	isValid, validCodFisc, err := utils.IsJWTTokenValid(token)

	var codFisc string = fmt.Sprintf("%v", validCodFisc)

	if err != nil {
		fmt.Println("[CHECKTOKEN, err]:", token, err)
		// http.Error(w, "Invalid Token", http.StatusForbidden)
		return isValid, codFisc, err
	}

	if !isValid {
		// w.Header().Set("Content-Type", "application/json")
		// http.Error(w, "Unauthorized", http.StatusUnauthorized)
		err := fmt.Errorf("Invalid Token")
		return isValid, codFisc, err
	}

	//str := "[CHECKTOKEN]: " + fmt.Sprintf("%v", codFisc)
	//fmt.Println(str)
	return isValid, codFisc, err
}

func getUserInfo(codFisc string) (*models.Utente, error) {
	db := myDBpckg.ConnectToDB()
	getUtente, err := db.Query("SELECT codFisc, nome, cognome, email FROM utente WHERE codFisc='" + codFisc + "';")
	if err != nil {
		return nil, err
	}

	utente := new(models.Utente)

	for getUtente.Next() {
		getUtente.Scan(&utente.CodFisc, &utente.Nome, &utente.Cognome, &utente.Email)
	}

	defer myDBpckg.CloseConnectionToDB(db)
	return utente, nil
}

func CheckHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {

		case "", "/": // ottengo il codFisc e mando al client tutte le sue info
			isValid, codFisc, err := checkRequest(w, r)
			if err != nil || !isValid {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			utente, _ := getUserInfo(codFisc)
			fmt.Println(utente)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(utente)

			next.ServeHTTP(w, r)

		case "/login", "/paziente/create", "/psicolgo/create":
			next.ServeHTTP(w, r)

		default:

			token := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
			if token == "" {
				http.Error(w, "Missing Token", http.StatusForbidden)
				return
			}

			isValid, codFisc, err := utils.IsJWTTokenValid(token)
			if err != nil {
				fmt.Println("[CHECKTOKEN, err]:", token, err)
				http.Error(w, "Invalid Token", http.StatusForbidden)
				return
			}

			if !isValid {
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			str := "[CHECKTOKEN]: " + fmt.Sprintf("%v", codFisc)
			fmt.Println(str)

			next.ServeHTTP(w, r)
		}
	})
}
