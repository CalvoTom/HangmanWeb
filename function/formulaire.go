package hangmanweb

import (
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

type Scoreboard struct {
	Username string
	Category string
	Points   int
}

type Account struct {
	Username string
	Email    string
	Password string
}

func Formulaire(w http.ResponseWriter, r *http.Request, Heasy *hangman.HangManData, Hmedium *hangman.HangManData, Hhard *hangman.HangManData, userscore *Scoreboard) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	log.Println(r.Form)
	for key, _ := range r.Form {
		if key == "playAgain" {
			playAgain(w, r, userscore)
		}
	}
	usernameCheckeur(w, r, userscore)
	gameCheckeur(w, r, Heasy, Hmedium, Hhard, userscore)

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func gameCheckeur(w http.ResponseWriter, r *http.Request, Heasy *hangman.HangManData, Hmedium *hangman.HangManData, Hhard *hangman.HangManData, userscore *Scoreboard) {
	input := r.FormValue("input")
	switch r.Header.Get("Referer") {
	case "http://localhost:8080/easy":
		hangman.Testeur(input, Heasy)
		userscore.Category = "Easy"
		CheckVictory(w, r, Heasy, userscore)
	case "http://localhost:8080/medium":
		hangman.Testeur(input, Hmedium)
		userscore.Category = "Medium"
		CheckVictory(w, r, Hmedium, userscore)
	case "http://localhost:8080/hard":
		hangman.Testeur(input, Hhard)
		userscore.Category = "Hard"
		CheckVictory(w, r, Hhard, userscore)
	}
}

func playAgain(w http.ResponseWriter, r *http.Request, userscore *Scoreboard) {
	switch userscore.Category {
	case "Easy":
		http.Redirect(w, r, "http://localhost:8080/estarting", http.StatusFound)
	case "Medium":
		http.Redirect(w, r, "http://localhost:8080/mstarting", http.StatusFound)
	case "Hard":
		http.Redirect(w, r, "http://localhost:8080/hstarting", http.StatusFound)
	}
}

func usernameCheckeur(w http.ResponseWriter, r *http.Request, userscore *Scoreboard) {
	username := r.FormValue("username")
	if username != "" {
		userscore.Username = username
		switch r.Header.Get("Referer") {
		case "http://localhost:8080/estarting":
			http.Redirect(w, r, "http://localhost:8080/easy", http.StatusFound)
		case "http://localhost:8080/mstarting":
			http.Redirect(w, r, "http://localhost:8080/medium", http.StatusFound)
		case "http://localhost:8080/hstarting":
			http.Redirect(w, r, "http://localhost:8080/hard", http.StatusFound)
		}
	}
}
