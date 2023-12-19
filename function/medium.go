package hangmanweb

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Medium(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/medium.html", "./templates/wordToFind.html", "./templates/information.html", "./templates/hangman.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}
