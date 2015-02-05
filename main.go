package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"

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
	router.HandleFunc("/api/books", ShowBooksAPI)

	// SERVER
	// ---------------
	fmt.Printf("The magic happens on port %s", port)
	http.ListenAndServe(":"+port, router)

}

// HANDLERS
// ---------------
func ShowBooks(w http.ResponseWriter, r *http.Request) {

	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	// First we parse the template.
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Then we execute it.
	// Note: You're executing the template and handling errors at the same time o.O
	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func ShowBooksAPI(w http.ResponseWriter, r *http.Request) {

	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	json, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
