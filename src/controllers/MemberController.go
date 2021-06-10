package controllers

import (
	"html/template"
	"log"
	"net/http"
)

var data Data

func HandleMembers(w http.ResponseWriter, r *http.Request) {
	log.Print("HandleMembers", r.Method, r.URL)
	switch r.Method {
	case "GET":
		handleGET(w, r)
	case "POST":
		handlePOST(w, r)
	}
}

func handleGET(w http.ResponseWriter, r *http.Request) {
	log.Print("handleGET")
	functions := template.FuncMap{"increment": func(i int) int { return i + 1 }}
	indexTmpl := template.Must(template.
		New("index.html").
		Funcs(functions).
		ParseFiles("src/static/index.html"))
	err := indexTmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}

func handlePOST(w http.ResponseWriter, r *http.Request) {
	log.Print("handlePOST ")
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	log.Print("name: ", name, " email: ", email)
	if data.Validate(name, email) && !data.MemberExists(email) {
		data.AddMember(name, email)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)	// redirect GET
}

