package hangmanweb

import (
	"net/http"

	hangman "github.com/CalvoTom/HangmanPackage"
)

func CheckVictory(w http.ResponseWriter, r *http.Request, H *hangman.HangManData, userscore *Scoreboard) {
	var tabscore []Scoreboard
	if H.Word == H.ToFind {
		userscore.Points = string(CalculateScore(H))
		tabscore = append(tabscore, *userscore)
		Save(tabscore)
		http.Redirect(w, r, "http://localhost:8080/victory", http.StatusFound)
	} else if H.Attempts == 0 {
		userscore.Points = "Loose"
		tabscore = append(tabscore, *userscore)
		Save(tabscore)
		http.Redirect(w, r, "http://localhost:8080/defeat", http.StatusFound)
	}
}
