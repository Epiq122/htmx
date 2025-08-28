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

		err := tmpl.ExecuteTemplate(w, "home.tmpl", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
