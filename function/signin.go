package hangmanweb

import (
	"html/template"
	"log"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/signin.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
