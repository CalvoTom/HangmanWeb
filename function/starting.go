package hangmanweb

import (
	hangman "github.com/CalvoTom/HangmanPackage"
)

func Starting(level string, Heasy *hangman.HangManData, Hmedium *hangman.HangManData, Hhard *hangman.HangManData) (He *hangman.HangManData, Hm *hangman.HangManData, Hh *hangman.HangManData) {
	Heasy = hangman.InitialiseStruc("words.txt")
	Hmedium = hangman.InitialiseStruc("words2.txt")
	Hhard = hangman.InitialiseStruc("words3.txt")

	return Heasy, Hmedium, Hhard
}
