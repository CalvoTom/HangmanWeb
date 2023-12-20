package hangmanweb

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request, data []Scoreboard) {
	template, err := template.ParseFiles("./index.html", "./templates/scoreboard.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, data)
}
