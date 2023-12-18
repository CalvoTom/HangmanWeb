package main

import (
	hangmanweb "hangmanweb/function"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func main() {
	H := hangman.InitialiseStruc("words.txt")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Home(w, r, H)
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Formulaire(w, r, H)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
