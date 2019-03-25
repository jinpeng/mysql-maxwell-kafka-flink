package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux" // specify used verbs
	"goapidemo/models"
	"goapidemo/controllers"
)

func defaultPage(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "Endpoint hit")
}

func handleRequests() {
	log.Println("Handling requests ...")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", defaultPage)
	router.HandleFunc("/api/1/contacts", controller.AllContacts).Methods("GET")
	router.HandleFunc("/api/1/contacts", controller.InsertContact).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	models.InitializeDatabase("contacts:password@tcp(database:3306)/contacts")
	handleRequests()
}