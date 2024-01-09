package hangmanweb

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Hard(w http.ResponseWriter, r *http.Request, Hhard *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/hard.html", "./templates/wordToFind.html", "./templates/information.html", "./templates/hangman.html", "./templates/lettersTried.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Hhard)
}
