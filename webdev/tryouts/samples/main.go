package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template.gohtml")
	if err != nil {
		fmt.Fprint(w, "<p style=\"color: red\">Page load failed!</p>")
		return
	}
	if err := t.Execute(w, struct {
		Name  string
		Color string
	}{
		Name:  "Asshole",
		Color: "coral",
	}); err != nil {
		fmt.Fprint(w, err)
	}
}

func mainStd() {
	http.HandleFunc("/", handleRoot) // htto.Handle("/", http.HandleFunc(handleRoot))
	http.ListenAndServe(":3000", nil)
}

func main() {
	router := chi.NewRouter()
	router.Get("/", handleRoot)
	fmt.Println("Listening @3000")
	http.ListenAndServe(":3000", router)
}
