package hangmanweb

import (
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func CheckVictory(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	if H.Word == H.ToFind {
		http.Redirect(w, r, "http://localhost:8080/victory", http.StatusFound)
	} else if H.Attempts == 0 {
		http.Redirect(w, r, "http://localhost:8080/defeat", http.StatusFound)
	}
}
