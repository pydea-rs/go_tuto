package main

import (
	htmlTemplatePackage "html/template"
	"log"
	"os"
	textTemplatePackage "text/template"
	// Another option is gobuffalo/plush
)

func xssSafe(data any) {
	tmp, err := htmlTemplatePackage.ParseFiles("./test.gohtml")
	if err != nil {
		log.Panicln(err)
	}

	if err = tmp.Execute(os.Stdin, data); err != nil {
		log.Panicln(err)
	}
}

func rawTemplate(data any) {
	tmp, err := textTemplatePackage.ParseFiles("./test.gohtml")
	if err != nil {
		log.Panicln(err)
	}

	if err = tmp.Execute(os.Stdin, data); err != nil {
		log.Panicln(err)
	}
}

func main() {

	data := struct {
		Name string
		Age  uint8
	}{
		Name: "<script>window.alert('Fucked up')</script>",
		Age:  28,
	}

	rawTemplate(data)
	log.Println("\n\nSafer Version with html/template:")
	xssSafe(data)
	log.Println("\n\nIf you want sth to not be encoded with html/template:")
	xssSafe(struct {
		Name htmlTemplatePackage.HTML // string that will not be enocded
		Age  uint8
	}{
		Name: `<script>window.alert('Fucked up with html/template')</script>`,
		Age:  29,
	})
}
