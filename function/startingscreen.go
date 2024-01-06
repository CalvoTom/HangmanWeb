package hangmanweb

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Estarting(w http.ResponseWriter, r *http.Request, Heasy *hangman.HangManData) *hangman.HangManData {
	template, err := template.ParseFiles("./pages/estarting.html")
	if err != nil {
		log.Fatal(err)
	}
	Heasy = hangman.InitialiseStruc("words.txt")
	template.Execute(w, nil)
	return Heasy
}

func Mstarting(w http.ResponseWriter, r *http.Request, Hmedium *hangman.HangManData) *hangman.HangManData {
	template, err := template.ParseFiles("./pages/mstarting.html")
	if err != nil {
		log.Fatal(err)
	}
	Hmedium = hangman.InitialiseStruc("words2.txt")
	template.Execute(w, nil)
	return Hmedium
}

func Hstarting(w http.ResponseWriter, r *http.Request, Hhard *hangman.HangManData) *hangman.HangManData {
	template, err := template.ParseFiles("./pages/hstarting.html")
	if err != nil {
		log.Fatal(err)
	}
	Hhard = hangman.InitialiseStruc("words3.txt")
	template.Execute(w, nil)
	return Hhard
}
