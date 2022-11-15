package main

import (
	"log"
	"net/http"

	"example.com/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)

}
