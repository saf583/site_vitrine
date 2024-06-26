package main

import (
	"fmt"
	"net/http"
	"projetweb/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.AccueilHandler).Methods("GET")
	router.HandleFunc("/achat", controllers.AchatHandler)
	router.HandleFunc("/contact", controllers.ContactHandler)
	router.HandleFunc("/evenement", controllers.EvenementHandler)
	router.HandleFunc("/compte", controllers.CompteHandler)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	fmt.Println("http://localhost:8081")
	http.ListenAndServe(":8081", router)
}
