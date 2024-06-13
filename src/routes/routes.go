package routes

import (
	"log"
	"net/http"

	"go/src/controllers"
	"github.com/gorilla/mux"
)

func HandlerRequest() {

	r := mux.NewRouter()

	r.HandleFunc("/amount", controllers.GetAmount).Methods("GET")
	r.HandleFunc("/createUser", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/sendMoney", controllers.TransferMoney).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
