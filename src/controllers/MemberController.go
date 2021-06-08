package controllers

import (
	"MemberClub/src/models"
	template2 "html/template"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Data struct {
	Members []models.Member
	Errors map[string]string
}

func (data *Data) Validate(name string, email string) bool {
	validateName := func(name string) bool {
		rgx := regexp.MustCompile("^([A-Za-z]+(\\s[a-zA-Z]+)*)$")
		return rgx.Match([]byte(name))
	}
	validateEmail := func(email string) bool {
		rgx := regexp.MustCompile(
			"^([A-Za-z0-9_\\.]+@[A-Za-z0-9]+\\.[A-Za-z0-9]+(\\.[A-Za-z0-9]+)*)$")
		return rgx.Match([]byte(email))
	}
	data.Errors = make(map[string]string)
	if !validateName(name) {
		data.Errors["name"] = "Name can contain only letters."
	}
	if !validateEmail(email) {
		data.Errors["email"] = "Enter valid email, please."
	}
	if len(data.Errors) == 0 {
		data.Members = append(data.Members, models.Member{
			Name:             name,
			Email:            email,
			RegistrationDate: time.Now(),
		})
	}
	return len(data.Errors) == 0
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
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		data.Validate(name, email)
		http.Redirect(w, r, "/", http.StatusSeeOther)	// redirect GET
	}
}
