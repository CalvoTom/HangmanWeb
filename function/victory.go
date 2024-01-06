package hangmanweb

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func Victory(w http.ResponseWriter, r *http.Request, userscore *Scoreboard, Heasy *hangman.HangManData, Hmedium *hangman.HangManData, Hhard *hangman.HangManData) {
	var tabscore []Scoreboard

	template, err := template.ParseFiles("./pages/victory.html")
	if err != nil {
		log.Fatal(err)
	}
	switch userscore.Category {
	case "Easy":
		userscore.Points = CalculateScore(Heasy)
	case "Medium":
		userscore.Points = CalculateScore(Hmedium)
	case "Hard":
		userscore.Points = CalculateScore(Hhard)
	}
	Save(tabscore)
	template.Execute(w, nil)
}
