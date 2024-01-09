package hangmanweb

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Easy(w http.ResponseWriter, r *http.Request, Heasy *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/easy.html", "./templates/wordToFind.html", "./templates/information.html", "./templates/hangman.html", "./templates/lettersTried.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Heasy)
}
