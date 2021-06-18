package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		var details ContactDetails
		details.Email = r.FormValue("email")
		details.Message = r.FormValue("message")
		details.Subject = r.FormValue("subject")

		fmt.Printf("%#v", details)
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
