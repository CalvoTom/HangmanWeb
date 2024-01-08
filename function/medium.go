package hangmanweb

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Medium(w http.ResponseWriter, r *http.Request, Hmedium *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/medium.html", "./templates/wordToFind.html", "./templates/information.html", "./templates/hangman.html", "./templates/lettersTried.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Hmedium)
}
