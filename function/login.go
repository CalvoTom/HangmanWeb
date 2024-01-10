package hangmanweb

import (
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/login.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
