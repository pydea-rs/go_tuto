package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	sample_input := `<script>window.alert("Youve been XSSed")</script>`
	safe_input := `&lt;script&gt;window.alert("XSS-Safe")&lt;/script&gt;`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, fmt.Sprint("<p>input="+sample_input+"</p>"))
	fmt.Fprint(w, fmt.Sprint("<p>input="+safe_input+"</p>"))

}

func main() {
	app := chi.NewRouter()
	app.Get("/", GetRoot)
	log.Println("Listening...")
	http.ListenAndServe(":8080", app)
}
