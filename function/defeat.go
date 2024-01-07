package hangmanweb

import (
	"html/template"
	"log"
	"net/http"
)

func Defeat(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/defeat.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
