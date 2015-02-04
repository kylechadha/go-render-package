package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {

	// CONFIG
	// --------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ROUTES
	// --------------
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", ShowBooks)

	// SERVER
	// ---------------
	fmt.Printf("The magic happens on port %s", port)
	http.ListenAndServe(":"+port, router)

}

// FUNCTIONS
// ---------------
func ShowBooks(w http.ResponseWriter, r *http.Request) {

	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
