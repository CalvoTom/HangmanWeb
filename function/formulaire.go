package hangmanweb

import (
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Formulaire(w http.ResponseWriter, r *http.Request, Heasy *hangman.HangManData, Hmedium *hangman.HangManData, Hhard *hangman.HangManData) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	input := r.FormValue("input")
	switch r.Header.Get("Referer") {
	case "http://localhost:8080/easy":
		hangman.Testeur(input, Heasy)
		log.Printf("easy")
	case "http://localhost:8080/medium":
		hangman.Testeur(input, Hmedium)
		log.Printf("medium")
	case "http://localhost:8080/hard":
		hangman.Testeur(input, Hhard)
		log.Printf("hard")
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
