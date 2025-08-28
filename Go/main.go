// working with a single template

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

// uppercase the string
func toUpper(str string) string {
	return strings.ToUpper(str)
}

// format the date
func formatDate(t time.Time) string {
	return t.Format("September 5, 1986")

}
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

	http.HandleFunc("/functions", func(w http.ResponseWriter, r *http.Request) {

		// our custom functions
		funcMap := template.FuncMap{
			"toUpper":    toUpper,
			"formatDate": formatDate,
		}

		tmpl := template.Must(template.New("functions.tmpl").Funcs(funcMap).ParseFiles("templates/functions.tmpl"))

		data := struct {
			Name        string
			CurrentDate time.Time
			Number      int
			Items       []string
		}{
			Name:        "Dr.Venkmin",
			CurrentDate: time.Now(),
			Number:      11,
			Items:       []string{"Apple", "Banana", "Cherry"},
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
