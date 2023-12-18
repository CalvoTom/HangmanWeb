package hangmanweb

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Home(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	template, err := template.ParseFiles("./index.html", "./templates/wordToFind.html", "./templates/information.html", "./templates/hangman.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}
