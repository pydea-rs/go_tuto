package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	text_template "text/template"

	"github.com/go-chi/chi/v5"
)

func GetRootRaw(w http.ResponseWriter, r *http.Request) {
	sample_input := `<script>window.alert("Youve been XSSed")</script>`
	safe_input := `&lt;script&gt;window.alert("XSS-Safe")&lt;/script&gt;`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, fmt.Sprint("<p>input="+sample_input+"</p>"))
	fmt.Fprint(w, fmt.Sprint("<p>input="+safe_input+"</p>"))

}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template.gohtml")

	if err != nil {
		fmt.Fprint(w, fmt.Sprint("Something went wrong: ", err))
		return
	}
	if err := t.Execute(w, struct {
		Bio  string
		Html template.HTML
	}{
		Bio:  `<script>window.alert("This one would be safe and escaped!")</script>`,
		Html: `<script>window.alert("This one would not be escaped")</script>`,
	}); err != nil {
		fmt.Fprint(w, fmt.Sprint("Cannot finalize the template: ", err))
		return
	}
}

func mainServer() {
	app := chi.NewRouter()
	app.Get("/", GetRoot)
	log.Println("Listening...")
	http.ListenAndServe(":8080", app)
}

func main() {
	fmt.Println("html/template Example:")
	t, err := template.ParseFiles("template.gohtml")
	if err != nil {
		panic(err)
	}

	if err := t.Execute(os.Stdout, struct {
		Bio  string
		Html template.HTML
	}{
		Bio:  `<script>window.alert("This one would be safe and escaped!")</script>`,
		Html: `<script>window.alert("This one would not be escaped")</script>`,
	}); err != nil {
		panic(err)
	}

	tx, err := text_template.ParseFiles("template.gohtml")
	if err != nil {
		panic(err)
	}

	fmt.Println("text/template Example:")
	if err = tx.Execute(os.Stdout, struct {
		Bio  string
		Html string // text template has no template.Html
	}{
		Bio:  `<script>window.alert("text template does not encode anything!")</script>`,
		Html: `<script>window.alert("text template does not encode anything")</script>`,
	}); err != nil {
		panic(err)
	}
}
