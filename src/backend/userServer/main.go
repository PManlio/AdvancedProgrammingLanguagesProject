package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./controllers"
	"./manager"
	"./middlewares"

	"github.com/gorilla/mux"
)

func homeRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Main Endpoint root /")
}

func handleRequests() {

	requestHandler := mux.NewRouter().StrictSlash(true)
	/*Il discorso consiste che questi SubRouter gestiscono delle, appunto,
	 *Sotto-route; sono a tutti gli effetti dei router che però SONO GENERATI
	 *dal router principale - qui definito come requestHandler. Questi sono dei
	 *puntatori al tipo Router e quindi possono essere mandati in input agli
	 *handler delle sotto-route.
	**/

	// RIMANE DA AGGIUNGERE IL MIDDLEWARE DELLE CORS POLICIES

	requestHandler.Use(middlewares.GlobalHeaders)
	requestHandler.Use(middlewares.CheckHeader)

	//momentaneo: Login
	loginSubrouter := requestHandler.PathPrefix("/login").Subrouter()

	pazienteSubrouter := requestHandler.PathPrefix("/paziente").Subrouter()
	psicologoSubrouter := requestHandler.PathPrefix("/psicologo").Subrouter()

	requestHandler.HandleFunc("/", homeRoot)

	manager.LoginHandler(loginSubrouter)
	controllers.PazientHandler(pazienteSubrouter)
	controllers.PsicologoHandler(psicologoSubrouter)

	// log.Fatal è l'equivalente di Print(), ma seguita da
	// una chiamata a os.Exit(1)
	log.Fatal(http.ListenAndServe(":8085", requestHandler))
}

func main() {
	fmt.Println("Server is running")

	// handleRequests() è "bloccante":
	// le funzioni successive non vengono eseguite
	handleRequests()

	// questa print, dopo handleRequests, non viene eseguita
	// fmt.Println("Server is running")
}
