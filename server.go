package main

import (
	hangmanweb "hangmanweb/function"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func main() {
	var userscore *hangmanweb.Scoreboard = new(hangmanweb.Scoreboard)
	var Heasy *hangman.HangManData
	var Hmedium *hangman.HangManData
	var Hhard *hangman.HangManData
	var level string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		scoreboardData := hangmanweb.OpenScoreboard()
		hangmanweb.Home(w, r, scoreboardData)
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Formulaire(w, r, Heasy, Hmedium, Hhard, userscore, level)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		level = r.URL.Query().Get("level")
		Heasy, Hmedium, Hhard = hangmanweb.Starting(level, Heasy, Hmedium, Hhard)
		hangmanweb.Login(w, r)
	})
	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		level = r.URL.Query().Get("level")
		Heasy, Hmedium, Hhard = hangmanweb.Starting(level, Heasy, Hmedium, Hhard)
		hangmanweb.Signin(w, r)
	})
	http.HandleFunc("/easy", func(w http.ResponseWriter, r *http.Request) {
		userscore.Category = "Easy"
		if userscore.Username == "" {
			http.Redirect(w, r, "http://localhost:8080/login?level=easy", http.StatusFound)
		}
		hangmanweb.Easy(w, r, Heasy)
	})
	http.HandleFunc("/medium", func(w http.ResponseWriter, r *http.Request) {
		userscore.Category = "Medium"
		if userscore.Username == "" {
			http.Redirect(w, r, "http://localhost:8080/login?level=medium", http.StatusFound)
		}
		hangmanweb.Medium(w, r, Hmedium)
	})
	http.HandleFunc("/hard", func(w http.ResponseWriter, r *http.Request) {
		userscore.Category = "Hard"
		if userscore.Username == "" {
			http.Redirect(w, r, "http://localhost:8080/login?level=hard", http.StatusFound)
		}
		hangmanweb.Hard(w, r, Hhard)
	})
	http.HandleFunc("/victory", func(w http.ResponseWriter, r *http.Request) {
		Heasy, Hmedium, Hhard = hangmanweb.Starting(level, Heasy, Hmedium, Hhard)
		hangmanweb.Victory(w, r)
	})
	http.HandleFunc("/defeat", func(w http.ResponseWriter, r *http.Request) {
		Heasy, Hmedium, Hhard = hangmanweb.Starting(level, Heasy, Hmedium, Hhard)
		hangmanweb.Defeat(w, r)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
