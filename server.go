package main

import (
	hangmanweb "hangmanweb/function"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func main() {
	var userscore *hangmanweb.Scoreboard = new(hangmanweb.Scoreboard)
	var Heasy *hangman.HangManData
	var Hmedium *hangman.HangManData
	var Hhard *hangman.HangManData

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		scoreboardData := hangmanweb.Open()
		log.Print()
		hangmanweb.Home(w, r, scoreboardData)
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Formulaire(w, r, Heasy, Hmedium, Hhard, userscore)
	})
	http.HandleFunc("/estarting", func(w http.ResponseWriter, r *http.Request) {
		Heasy = hangmanweb.Estarting(w, r, Heasy)
	})
	http.HandleFunc("/mstarting", func(w http.ResponseWriter, r *http.Request) {
		Hmedium = hangmanweb.Mstarting(w, r, Hmedium)
	})
	http.HandleFunc("/hstarting", func(w http.ResponseWriter, r *http.Request) {
		Hhard = hangmanweb.Hstarting(w, r, Hmedium)
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
	http.HandleFunc("/victory", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Victory(w, r)
	})
	http.HandleFunc("/defeat", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Defeat(w, r)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
