package hangmanweb

import (
	"html/template"
	"log"
	"net/http"
)

func Victory(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/victory.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
