package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"../utils"
)

func CheckHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {

		case "/login", "/paziente/create", "/psicolgo/create":
			next.ServeHTTP(w, r)

		default:

			token := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
			if token == "" {
				http.Error(w, "Missing Token", http.StatusForbidden)
				return
			}

			isValid, err := utils.IsJWTTokenValid(token)
			if err != nil {
				fmt.Println(token)
				http.Error(w, "Invalid Token", http.StatusForbidden)
				return
			}

			if !isValid {
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			next.ServeHTTP(w, r)
		}
	})
}
