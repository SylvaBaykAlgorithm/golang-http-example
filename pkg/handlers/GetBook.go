package handlers

import (
	"net/http"
	"strconv"

	"encoding/json"

	"example.com/pkg/mocks"
	"github.com/gorilla/mux"
)

func getBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	for _, book := range mocks.Books {
		if book.Id == id {
			// If ids are equal send book as a response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
