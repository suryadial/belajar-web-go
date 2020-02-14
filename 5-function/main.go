package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type SuperHero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s SuperHero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = SuperHero{
			Name:    "Bruce Wayne",
			Alias:   "Batman",
			Friends: []string{"Superman", "Flash", "Green Lantern"},
		}

		var tmpl = template.Must(template.ParseFiles("views.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("server start at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
