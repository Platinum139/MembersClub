package controllers

import (
	"MemberClub/src/models"
	template2 "html/template"
	"log"
	"net/http"
	"time"
)

type Data struct {
	Items []models.Member
}

var data Data

func HandleMembers(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	if r.Method == "GET" {
		functions := template2.FuncMap{"increment": func(i int) int { return i + 1 }}
		template := template2.Must(template2.
				New("index.html").
				Funcs(functions).
				ParseFiles("src/static/index.html"))
		err := template.Execute(w, data)
		if err != nil {
			log.Println(err)
			return
		}
	}
	if r.Method == "POST" {
		// get form data
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")

		// create member
		registrationDate := time.Now()
		newMember := models.Member{
			Name: name,
			Email: email,
			RegistrationDate: registrationDate,
		}
		// add new member to array
		data.Items = append(data.Items, newMember)

		// redirect with method GET
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
