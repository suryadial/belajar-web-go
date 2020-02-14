package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (i Info) GetAffiliationDetailInfo() string {
	return "Have 31 divisions"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Suryadi Misla",
			Gender:  "Man",
			Hobbies: []string{"Reading", "Coding", "Sleeping"},
			Info:    Info{"Wayne Enterprise", "Gotham City"},
		}

		var tmpl = template.Must(template.ParseFiles("views.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("server start at localhost:9000")
	http.ListenAndServe(":9000", nil)

}
