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

func Formulaire(w http.ResponseWriter, r *http.Request, Heasy *hangman.HangManData, Hmedium *hangman.HangManData, Hhard *hangman.HangManData, userscore *Scoreboard, level string) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	for key, _ := range r.Form {
		if key == "playAgain" {
			playAgain(w, r, level)

		} else if key == "signin" {
			http.Redirect(w, r, "http://localhost:8080/signin?level="+level, http.StatusFound)
		} else if key == "login" {
			http.Redirect(w, r, "http://localhost:8080/login?level="+level, http.StatusFound)
		} else if key == "username-sign" {
			newUsername := r.FormValue("username-sign")
			newEmail := r.FormValue("email-sign")
			newPassword := r.FormValue("password-sign")

			if singIn(w, r, newUsername, newEmail, userscore) {
				var listeNewAccount []Account
				userscore.Username = newUsername
				listeNewAccount = append(listeNewAccount, Account{Username: newUsername, Email: newEmail, Password: newPassword})

				SaveAccount(listeNewAccount)
				http.Redirect(w, r, "http://localhost:8080/"+level, http.StatusFound)
			} else {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			}

		} else if key == "email-log" {
			email := r.FormValue("email-log")
			password := r.FormValue("password-log")

			if logIn(w, r, email, password, userscore) {
				log.Println(level)
				http.Redirect(w, r, "http://localhost:8080/"+level, http.StatusFound)
			} else {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			}

		} else {
			gameCheckeur(w, r, Heasy, Hmedium, Hhard, userscore)
		}
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func gameCheckeur(w http.ResponseWriter, r *http.Request, Heasy *hangman.HangManData, Hmedium *hangman.HangManData, Hhard *hangman.HangManData, userscore *Scoreboard) {
	input := r.FormValue("input")
	switch r.Header.Get("Referer") {
	case "http://localhost:8080/easy":
		hangman.Testeur(input, Heasy)
		CheckVictory(w, r, Heasy, userscore)
	case "http://localhost:8080/medium":
		hangman.Testeur(input, Hmedium)
		CheckVictory(w, r, Hmedium, userscore)
	case "http://localhost:8080/hard":
		hangman.Testeur(input, Hhard)
		CheckVictory(w, r, Hhard, userscore)
	}
}

func playAgain(w http.ResponseWriter, r *http.Request, level string) {
	http.Redirect(w, r, "http://localhost:8080/"+level, http.StatusFound)
}

func logIn(w http.ResponseWriter, r *http.Request, email string, password string, userscore *Scoreboard) bool {
	listeAccount := OpenAccount()

	for i := 0; i < len(listeAccount); i++ {
		if listeAccount[i].Email == email && listeAccount[i].Password == password {
			userscore.Username = listeAccount[i].Username
			return true
		}
	}
	return false
}

func singIn(w http.ResponseWriter, r *http.Request, newUsername string, newEmail string, userscore *Scoreboard) bool {
	listeAccount := OpenAccount()
	for i := 0; i < len(listeAccount); i++ {
		if listeAccount[i].Email == newEmail {
			return false
		} else if listeAccount[i].Username == newUsername {
			return false
		}
	}
	return true
}
