package hangmanweb

import (
	"fmt"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Formulaire(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		input := r.FormValue("input")
		hangman.Testeur(input, H)

		http.Redirect(w, r, "http://localhost:8080", http.StatusFound)

	default:
		fmt.Fprintf(w, "Only POST")
	}
}
