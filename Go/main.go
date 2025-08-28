// working with a single template

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseGlob("templates/*.tmpl"))

		data := struct {
			Name        string
			Title       string
			Description string
			Socials     map[string]string
			Features    []string
		}{
			Name:        "Dr.Venkmin",
			Title:       "Visitor",
			Description: "Welcome to the Go web development slammmm piece.",
			Socials: map[string]string{
				"GitHub":   "https://github.com ",
				"Twitter":  "https://twitter.com/yourprofile",
				"LinkedIn": "https://linkedin.com/in/yourprofile",
			},
			Features: []string{
				"Fast and Efficient",
				"Scalable",
				"Easy to Learn",
			},
		}

		err := tmpl.ExecuteTemplate(w, "home.tmpl", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
