package hangmanweb

import hangman "github.com/CalvoTom/HangmanPackage"

func CalculateScore(H *hangman.HangManData) int {
	return H.Attempts * 10
}
