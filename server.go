package main

import (
	hangmanweb "hangmanweb/function"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func main() {
	var userscore *hangmanweb.Scoreboard = new(hangmanweb.Scoreboard)
	scoreboardData := hangmanweb.Open()
	Heasy := hangman.InitialiseStruc("words.txt")
	Hmedium := hangman.InitialiseStruc("words2.txt")
	Hhard := hangman.InitialiseStruc("words3.txt")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Home(w, r, scoreboardData)
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Formulaire(w, r, Heasy, Hmedium, Hhard, userscore)
	})
	http.HandleFunc("/easy", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Easy(w, r, Heasy)
	})
	http.HandleFunc("/medium", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Medium(w, r, Hmedium)
	})
	http.HandleFunc("/hard", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Hard(w, r, Hhard)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
