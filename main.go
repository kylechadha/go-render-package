package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Render *render.Render

func main() {

	// CONFIG
	// --------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	Render = render.New(render.Options{})

	// ROUTES
	// --------------
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", ShowBooks)             // vanilla html route
	router.HandleFunc("/api/books", ShowBooksAPI) // vanilla json route
	router.HandleFunc("/render/data", DataRender) // render pkg routes
	router.HandleFunc("/render/json", JsonRender) // render pkg routes
	router.HandleFunc("/render/html", HtmlRender) // render pkg routes

	// SERVER
	// ---------------
	fmt.Printf("The magic happens on port %s", port)
	http.ListenAndServe(":"+port, router)

}

// HANDLERS
// ---------------
func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Kyle Chadha"}

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
	book := Book{"Building Web Apps with Go", "Kyle Chadha"}

	json, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func DataRender(w http.ResponseWriter, r *http.Request) {
	Render.Data(w, http.StatusOK, []byte("Some binary data here."))
}

func JsonRender(w http.ResponseWriter, r *http.Request) {
	Render.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
}

func HtmlRender(w http.ResponseWriter, r *http.Request) {
	Render.HTML(w, http.StatusOK, "example", string("this is a string"))
}
