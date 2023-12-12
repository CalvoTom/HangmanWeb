package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func formulaire(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		input := r.FormValue("input")
		log.Println(input)
		http.Redirect(w, r, "http://localhost:8080", http.StatusFound)

	default:
		fmt.Fprintf(w, "Only POST")
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", formulaire)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
