package hangmanweb

import (
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

type Scoreboard struct {
	Username string
	Points   int
}

func Formulaire(w http.ResponseWriter, r *http.Request, Heasy *hangman.HangManData, Hmedium *hangman.HangManData, Hhard *hangman.HangManData, userscore *Scoreboard) {
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

	var tabscore []Scoreboard
	if len(r.FormValue("username")) != 0 {
		userscore.Username = r.FormValue("username")
		userscore.Points = 100
		tabscore = append(tabscore, *userscore)
		Save(tabscore)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
